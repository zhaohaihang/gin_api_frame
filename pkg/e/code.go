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

	// 活动错误

	// token 错误
	ErrorAuthCheckTokenFail        = 30001 
	ErrorAuthCheckTokenTimeout     = 30002 
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
