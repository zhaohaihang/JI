package e

const (
	// HTTP code
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400

	// 用户错误
	ErrorNotCompare         = 10004

	// 活动错误

	// token 错误
	ErrorAuthCheckTokenFail        = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     = 30002 //token 过期
	ErrorAuthToken                 = 30003
	ErrorAuthInsufficientAuthority = 30005

	// //数据库错误
	ErrorDatabase = 40001


)
