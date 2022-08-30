package resp

import (
	"resume-server/utils"
	"time"
)

// ---------------------------------------------
// 返回状态码
// 2 开头为成功，1 开头为常见错误，4 开头为其他错误
// ---------------------------------------------

var (
	EthIsOK = newBack(2000, "请求成功")

	ErrServer = newBack(1001, "系统服务异常")
	ErrParam  = newBack(1002, "参数有误，请检查")

	RegisterUserNameErr  = newBack(4001, "注册用户名已存在")
	RegisterUserEmailErr = newBack(4002, "注册邮箱已存在")
	RegisterOTPExpired   = newBack(4003, "验证码错误或已过期")

	LoginFail         = newBack(4000, "邮箱或用户名不能同时为空")
	LoginUserNameErr  = newBack(4004, "用户名或密码错误")
	LoginUserEmailErr = newBack(4005, "邮箱或密码错误")

	AuthFail       = newBack(4010, "用户认证失败")
	AuthSendOTPErr = newBack(4011, "发送验证码邮件失败")
	AuthExpired    = newBack(4012, "用户认证信息过期")

	ResumeErr = newBack(4020, "该条简历不存在")
)

// ---------------------------------------------
// 返回方法
// ---------------------------------------------

// 返回内容
type back struct {
	Code int         `json:"code"`          // 业务编码
	Msg  string      `json:"msg"`           // 错误描述
	Data interface{} `json:"data"`          // 成功时返回的数据
	QID  string      `json:"qid,omitempty"` // 当前请求的唯一ID，便于问题定位，忽略也可以
}

// 创建返回参数
func newBack(code int, msg string) (b *back) {
	return &back{
		Code: code,
		Msg:  msg,
	}
}

// With 调用该方法进行返回
func With(reply *back, content any) (data any) {
	reply.Data = content
	reply.QID = time.Now().Format("060102150405") + utils.RandLow(6) // 生成QID的方法，到秒的时间+6位随机字符串
	return reply
}
