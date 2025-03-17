package e

var MsgFlags = map[int]string{
	Success: "ok",
	Error: "fail",
	InvalidParams: "参数错误",
	ErrorExitUser: "用户名已存在",
	ErrorFailEncryption: "密码加密失败",		
	ErrorExistUserNotFound: "用户不存在",	
	ErrorNotCompare: "密码错误",
	ErrorAuthToken: "token认证失败",
	ErrorAuthCheckTokenTimeout: "token过期",
	ErrorUploadFail: "图片上传失败",
	ErrorSendEmail: "邮件发送失败",
	ErrorProductImgUpload: "图片上传错误",
	ErrorDatabase: "数据库错误",
}

//getmsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}