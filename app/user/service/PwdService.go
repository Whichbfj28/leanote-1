package service

import (
	"context"

	"github.com/coocn-cn/leanote/app/db"
	"github.com/coocn-cn/leanote/app/info"
	. "github.com/coocn-cn/leanote/app/lea"
	"gopkg.in/mgo.v2/bson"
)

// 找回密码
// 修改密码
var overHours = 2.0 // 小时后过期

type PwdService struct {
	userSrv  *UserService
	tokenSrv *TokenService
	emailSrv EmailService
}

func NewPWD(ctx context.Context, userSrv *UserService, tokenSrv *TokenService, emailSrv EmailService) *PwdService {
	return &PwdService{
		userSrv:  userSrv,
		tokenSrv: tokenSrv,
		emailSrv: emailSrv,
	}
}

// 1. 找回密码, 通过email找用户,
// 用户存在, 生成code
func (m *PwdService) FindPwd(email string) (ok bool, msg string) {
	ok = false
	userId := m.userSrv.GetUserId(email)
	if userId == "" {
		msg = "用户不存在"
		return
	}

	token := m.tokenSrv.NewToken(userId, email, info.TokenPwd)
	if token == "" {
		return false, "db error"
	}

	// 发送邮件
	ok, msg = m.emailSrv.FindPwdSendEmail(token, email)
	return
}

// 重置密码时
// 修改密码
// 先验证
func (m *PwdService) UpdatePwd(token, pwd string) (bool, string) {
	var tokenInfo info.Token
	var ok bool
	var msg string

	// 先验证
	if ok, msg, tokenInfo = m.tokenSrv.VerifyToken(token, info.TokenPwd); !ok {
		return ok, msg
	}

	passwd := GenPwd(pwd)
	if passwd == "" {
		return false, "GenerateHash error"
	}

	// 修改密码之
	ok = db.UpdateByQField(db.Users, bson.M{"_id": tokenInfo.UserId}, "Pwd", passwd)

	// 删除token
	m.tokenSrv.DeleteToken(tokenInfo.UserId.Hex(), info.TokenPwd)

	return ok, ""
}
