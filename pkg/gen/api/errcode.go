package api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type ErrCode interface {
	Code() int
	String() string
}

type Err struct {
	code   int
	msg    string
	detail error
}

func (e Err) Code() int {
	return e.code
}

func (e Err) String() string {
	if gin.Mode() == gin.ReleaseMode {
		return e.msg
	}
	if e.detail != nil {
		return e.msg + ", detail: " + e.detail.Error()
	}
	return e.msg
}

func (e Err) Wrap(err interface{}) ErrCode {
	switch err.(type) {
	case error:
		e.detail = err.(error)
	case string:
		e.detail = errors.New(err.(string))
	}
	return e
}

func NewErr(e Err, err interface{}) Err {
	switch err.(type) {
	case error:
		e.detail = err.(error)
	case string:
		e.detail = errors.New(err.(string))
	}
	return e
}

func Equal(a, b ErrCode) bool {
	return a.Code() == b.Code()
}

var (
	ECSuccess           = Err{code: 0, msg: "success"}
	ECInternal          = Err{code: 500, msg: "服务器内部错误"}
	ECDbFind            = Err{code: 1000, msg: "数据库查询错误"}
	ECDbFirst           = Err{code: 1001, msg: "数据库查找错误"}
	ECDbSave            = Err{code: 1002, msg: "数据库保存错误"}
	ECDbCreate          = Err{code: 1003, msg: "数据库创建错误"}
	ECDbDelete          = Err{code: 1003, msg: "数据库删除错误"}
	ECDbExec            = Err{code: 1004, msg: "数据库操作错误"}
	ECDbCommit          = Err{code: 1005, msg: "数据库事务提交错误"}
	ECParam             = Err{code: 2000, msg: "参数错误"}
	ECPhoneNumber       = Err{code: 2001, msg: "手机错误"}
	ECPhoneCode         = Err{code: 2002, msg: "手机验证码无效"}
	ECParamRefreshToken = Err{code: 2003, msg: "刷新token错误"}
	ECSms               = Err{code: 3002, msg: "手机验证码发送失败"}
	ECTokenGen          = Err{code: 3003, msg: "token生成错误"}
)
