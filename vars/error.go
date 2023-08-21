package vars

type CodeError struct {
	Code    int
	Message string
}

func (e CodeError) Error() string {
	return e.Message
}

// type ErrCode int

const (
	// 通用错误
	ErrUserNotFound        = -1000
	ErrUserExist           = -1001
	ErrUserInvalidPassword = -1002

	// 业务错误
	ErrImgUploadFail = -110100

	// 数据库相关

)

var ErrorMap = map[int]CodeError{
	ErrUserExist:           {ErrUserExist, "用户已存在"},
	ErrUserNotFound:        {ErrUserNotFound, "用户不存在"},
	ErrUserInvalidPassword: {ErrUserInvalidPassword, "密码错误"},
	ErrImgUploadFail:       {ErrImgUploadFail, "图片上传失败"},
}
