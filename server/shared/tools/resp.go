package tools

import (
	"errors"
	"project/server/shared/errno"
	"project/server/shared/kitex_gen/base"
)

func BuildBaseResp(err error) *base.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *base.BaseResponse {
	return &base.BaseResponse{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}

func ParseBaseResp(resp *base.BaseResponse) error {
	if resp.StatusCode == errno.Success.ErrCode {
		return nil
	}
	return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
}
