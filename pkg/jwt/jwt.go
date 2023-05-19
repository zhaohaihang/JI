package jwt

import (
	"fmt"
	"ji/pkg/e"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"time"
)


var jwtSecret = []byte("FanOne")

type Claims struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

//GenerateToken 签发用户Token
func GenerateToken(id uint, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserID:    id,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ji",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//ParseToken 验证用户token
func ParseToken(tokenStr string) (*Claims, error) {
	token := strings.Fields(tokenStr)
	if len(token) != 2 || strings.ToLower(token[0]) != "Bearer" || token[1] == "" {
		return nil, fmt.Errorf("authorization header invaild")
	}

	tokenClaims, err := jwt.ParseWithClaims(token[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func SetTokenClaimsToContext(c *gin.Context, claims *Claims) {
	if c == nil || claims == nil {
		return
	}
	c.Set("claims", claims)
}

func GetTokenClaimsFromContext(c *gin.Context) *Claims {
	if c == nil {
		return nil
	}

	val, ok := c.Get("claims")
	if !ok {
		return nil
	}

	claims, ok := val.(*Claims)
	if !ok {
		return nil
	}

	return claims
}

//JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		var claims *Claims
		if token == "" {
			code = e.ErrorTokenIsNUll
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		SetTokenClaimsToContext(c, claims)
		c.Next()
	}
}