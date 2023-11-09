package main

import (
	"context"
	"project/server/cmd/user/pkg/mysql"
	"project/server/shared/consts"
	"project/server/shared/errno"
	"project/server/shared/kitex_gen/user"
	"project/server/shared/tools"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/hertz-contrib/paseto"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	UserMysqlManager
	TokenGenerator
	RedisManager
}

type UserMysqlManager interface {
	CreateUser(ctx context.Context, user *mysql.User) error
	ExistOrNotByMobile(ctx context.Context, mobile string) (user *mysql.User, exist bool, err error)
}

type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

type RedisManager interface {
	SavePhoneMsg(ctx context.Context, phone, code string) error
	GetPhoneMsg(ctx context.Context, phone string) (code string, err error)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// PhoneLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) PhoneLogin(ctx context.Context, req *user.PhoneLoginRequest) (resp *user.PhoneLoginResponse, err error) {
	// TODO: Your code here...
	code, err := s.GetPhoneMsg(ctx, req.Mobile)
	if err != nil {
		klog.Error("user phone verify code err: ", err)
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr.WithMessage("phone verify code err"))
		return resp, nil
	}
	if code != req.Code {
		klog.Error("user phone verify code err: ", err)
		resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr.WithMessage("phone verify code err"))
		return resp, nil
	}

	user, exist, err := s.ExistOrNotByMobile(ctx, req.Mobile)
	if !exist {
		userId, _ := snowflake.NewNode(consts.UserSnowflakeNode)
		nickName := tools.GenerateRandomNickNameString()
		status := mysql.Active
		u := &mysql.User{
			Mobile:   req.Mobile,
			NickName: nickName,
			UserId:   userId.Generate().Int64(),
			Status:   status,
		}
		if err := s.CreateUser(ctx, u); err != nil {
			klog.Error("create user err:", err)
			resp.BaseResp = tools.BuildBaseResp(errno.UserSrvErr)
			return resp, nil
		}
	}

	now := time.Now()
	resp.Token, err = s.TokenGenerator.CreateToken(&paseto.StandardClaims{
		ID:        strconv.FormatInt(user.UserId, 10),
		Issuer:    consts.Issuer,
		Audience:  consts.User,
		IssuedAt:  now,
		NotBefore: now,
		ExpiredAt: now.Add(consts.ThirtyDays),
	})
	if err != nil {
		klog.Error("create token error:", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		return resp, nil
	}

	resp.BaseResp = tools.BuildBaseResp(nil)
	return
}
