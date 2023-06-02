package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	UpdatePasswordSuccess: "修改密码成功",
	NotExistInentifier:    "该第三方账号未绑定",
	ERROR:                 "fail",
	InvalidParams:         "请求参数错误",

	ErrorPasswordNotCompare: "账号密码错误",
	ErrorUserCreate:         "创建用户错误",

	ErrorAuthCheckTokenFail:        "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:     "Token已超时",
	ErrorAuthToken:                 "Token生成失败",
	ErrorAuthInsufficientAuthority: "权限不足",
	ErrorTokenIsNUll:               "未携带token数据",

	ErrorDatabase: "数据库操作出错,请重试",

	ErrorUploadFile:        "文件上传失败",
	ErrorUserActivityLimit: "创建活动太频繁",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
