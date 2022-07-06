package errmsg


const (
	SUCCSE = 200
	ERROR = 500
	// code =1000 ... 用户模块的错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIT = 1003
	ERROR_TOKEN_EXIT = 1004
	ERROR_TOKEN_RUNTIME = 1005
	ERROR_TOKEN_WRONG = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code = 2000 ... 文章模块的错误
	// code = 3000 ... 分类模块的错误

)

var codemsg = map[int]string{
	SUCCSE: "ok",
	ERROR: "FAIL",
	ERROR_USERNAME_USED:"用户名已存在",
	ERROR_PASSWORD_WRONG:"密码错误",
	ERROR_USER_NOT_EXIT:"用户不存在",
	ERROR_TOKEN_EXIT:"TOKEN不存在",
	ERROR_TOKEN_RUNTIME:"TOKEN不存在",
	ERROR_TOKEN_WRONG:"TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:"TOKEN格式错误",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}