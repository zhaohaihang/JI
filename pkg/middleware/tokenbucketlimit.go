package middleware

import (
	"fmt"
	"ji/pkg/e"
	"ji/pkg/redis"
	"ji/pkg/utils/tokenutil.go"
	"time"

	"github.com/gin-gonic/gin"
	redigo "github.com/gomodule/redigo/redis"
)

const (
	KeyTokenBucketLimitActivityUser = "tokenbucketlimit:activity:username:%s"
)

func BucketLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := e.SUCCESS

		claims := tokenutil.GetTokenClaimsFromContext(c)
		if claims == nil {
			code = e.ErrorUserActivityLimit
			// c.JSON(http.StatusBadRequest, gin.H{"message": "username获取失败"})
		} else {
			username := claims.Username
			key := fmt.Sprintf(KeyTokenBucketLimitActivityUser, username)
			conn := redis.Pool.Get() //从连接池，取一个链接
			defer conn.Close()
			rate := 1                                                     // 令牌生成速度 每秒1个token
			capacity := 1                                                 // 桶容量
			tokens, _ := redigo.Int(conn.Do("hget", key, "tokens"))       // 桶中的令牌数
			lastTime, _ := redigo.Int64(conn.Do("hget", key, "lastTime")) // 上次令牌生成时间
			now := time.Now().Unix()
			// 初始状态下 令牌数量为桶的容量
			existKey, _ := redigo.Int(conn.Do("exists", key))
			if existKey != 1 {
				tokens = capacity
				conn.Do("hset", key, "lastTime", now)
			}
			deltaTokens := int(now-lastTime) * rate // 经过一段时间后生成的令牌
			if deltaTokens > 1 {
				tokens = tokens + deltaTokens // 增加令牌
			}
			if tokens < 1 {
				code = e.ErrorUserActivityLimit
			} else {
				if tokens > capacity {
					tokens = capacity
				}
				tokens-- // 请求进来了，令牌就减少1
				conn.Do("hset", key, "lastTime", now)
				conn.Do("hset", key, "tokens", tokens)
				c.Next()
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
		c.Next()
	}
}
