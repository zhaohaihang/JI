package service

type LoginUserInfo struct {
	UserName  string `form:"user_name" json:"user_name"`
	Password  string `form:"password" json:"password"`
    Type      int    `form:"type" json:"type"` 
}