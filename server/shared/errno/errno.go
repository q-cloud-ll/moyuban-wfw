package errno

import (
	"fmt"
	"project/server/shared/kitex_gen/errno"
)

var (
	Success        = NewErrNo(int64(errno.Err_Success), "success")
	NoRoute        = NewErrNo(int64(errno.Err_NoRoute), "no route")
	NoMethod       = NewErrNo(int64(errno.Err_NoMethod), "no method")
	BadRequest     = NewErrNo(int64(errno.Err_BadRequest), "bad request")
	ParamsErr      = NewErrNo(int64(errno.Err_ParamsErr), "params error")
	AuthorizeFail  = NewErrNo(int64(errno.Err_AuthorizeFail), "authorize failed")
	TooManyReqeust = NewErrNo(int64(errno.Err_TooManyRequest), "too many requests")
	ServiceErr     = NewErrNo(int64(errno.Err_ServiceErr), "service error")
	RPCUserSrvErr  = NewErrNo(int64(errno.Err_RPCUserSrvErr), "rpc user service error")
	UserSrvErr     = NewErrNo(int64(errno.Err_UserSrvErr), "user service error")
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// NewErrNo return ErrNo
func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}
