// Code generated by hertz generator.

package user

import (
	"context"
	"net/http"
	"project/server/cmd/api/config"
	"project/server/shared/errno"
	"project/server/shared/tools"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"project/server/cmd/api/biz/model/user"
	kuser "project/server/shared/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PhoneSendMsg .
// @router /user/sendMsg [POST]
func PhoneSendMsg(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.PhoneSendMsgRequest
	resp := new(kuser.PhoneSendVerifyCodeResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	res, err := config.GlobalUserClient.PhoneSendVerifyCode(ctx, &kuser.PhoneSendVerifyCodeRequest{
		Mobile: req.Mobile,
	})
	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, res)
}

// Captcha .
// @router /user/captcha/request [GET]
func Captcha(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CaptchaRequest
	resp := new(kuser.CaptchaResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	res, err := config.GlobalUserClient.Captcha(ctx, &kuser.CaptchaResquest{
		CaptchaId: req.CaptchaId,
	})

	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, res)
}

// PhoneLogin .
// @router /user/phoneLogin [POST]
func PhoneLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.PhoneLoginRequest
	resp := new(kuser.PhoneLoginResponse)
	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	res, err := config.GlobalUserClient.PhoneLogin(ctx, &kuser.PhoneLoginRequest{
		Mobile: req.Mobile,
		Code:   req.Code,
	})
	if err != nil {
		hlog.Error("rpc user service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(consts.StatusOK, res)
}
