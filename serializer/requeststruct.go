package serializer

type LoginUserInfo struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,emailorphone"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
	Type     int    `form:"type" json:"type" binding:"required,oneof=1 2"`
}

type Point struct {
	Lat float64 `form:"lat" json:"lat" binding:"latitude"`
	Lng float64 `form:"lng" json:"lng" binding:"longitude"`
}
type UpdateUserInfo struct {
	Biography string `form:"biography" json:"biography" binding:"max=1000"`
	Address   string `form:"address" json:"address" binding:"max=1000"`
	Email     string `form:"email" json:"email" binding:"email"`
	Phone     string `form:"phone" json:"phone" binding:"phone"`
	Location  Point  `form:"location" json:"location" binding:"max=100"`
	Extra     string `form:"extra" json:"extra" binding:"max=1000"`
}