package e 

const (
	Success = 200
	Error = 500
	InvalidParams = 400
	//user模块错误
	ErrorExitUser = 30001
	ErrorFailEncryption = 30002
	ErrorExistUserNotFound = 30003
	ErrorNotCompare = 30004
	ErrorAuthToken = 30005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUploadFail = 30007
	ErrorSendEmail = 30008
	//product 模块 4xxxx
	ErrorProductImgUpload = 40001
	ErrorDatabase = 40002

	//redis
	ErrorRedis = 50001

	ErrorProductMoreCart = 20008
	ErrorProductExistCart = 20009
)