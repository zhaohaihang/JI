package e

const (
	// HTTP code
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400

	// 用户错误
	ErrorPasswordNotCompare = 10001
	ErrorUserCreate = 10002
	ErrorGetUserInfo = 10003

	// 活动错误
	ErrorActivityCreate = 20001

	// token 错误
	ErrorAuthCheckTokenFail        = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     = 30002 //token 过期
	ErrorAuthToken                 = 30003
	ErrorTokenIsNUll               = 30004
	ErrorAuthInsufficientAuthority = 30005

	// //数据库错误
	ErrorDatabase = 40001

	// 静态资源错误
	ErrorUploadFile = 50001

	// 限流
	ErrorUserActivityLimit = 60001
)
