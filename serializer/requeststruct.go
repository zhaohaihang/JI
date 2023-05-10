package serializer

type LoginUserInfo struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,emailorphone"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
	Type     int    `form:"type" json:"type" binding:"required,oneof=1 2"`
}
