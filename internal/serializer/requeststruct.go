package serializer

type BasePage struct {
	PageNum  int `form:"page_num" json:"page_num"`
	PageSize int `form:"page_size" json:"page_size"`
}

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
	Location  Point  `form:"location" json:"location"`
	Extra     string `form:"extra" json:"extra" binding:"max=1000"`
}

type CreateActivityInfo struct {
	Title          string `form:"title" json:"title" binding:"max=30"`
	Introduction   string `form:"introduction" json:"introduction" binding:"max=1000"`
	Status         int    `form:"status" json:"status" binding:"required,oneof=1 2 3"`
	StartTime      int64  `form:"start_time" json:"start_time"  binding:"required"`
	EndTime        int64  `form:"end_time" json:"end_time"  binding:"required"`
	Location       Point  `form:"location" json:"location"`
	ExpectedNumber uint   `form:"expected_number" json:"expected_number" `
}

type NearInfo struct {
	Point
	Rad int `form:"rad" json:"rad" binding:"lte=100"`
}
