package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

/*
	注册：用户使用邮箱注册（要求：账号为邮箱账号，密码为6-12位数字字母组合）
	流程为：用户使用邮箱注册，填写邮箱账号，验证邮箱有效性，有效注册成功，无效或限定时间内未进行验证为无效注册，删除该用户条目

	后台逻辑为：读取邮箱参数，进行邮箱校验，失败返回邮箱格式错误信息，
	读取密码参数,进行密码正则校验，校验失败返回密码格式错误，
	进行数据库唯一校验，校验唯一存储账号状态为未激活状态，
	发送验证邮件，验证后进行状态更新，将未激活账户改为已激活账户。

*/

var (
	w       *httptest.ResponseRecorder
	c       *gin.Context
	control UserController
)

func setupSignup() {
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/hola", nil)
	c.Params = gin.Params{gin.Param{Key: "account", Value: "1"}}
	c.Set("account", "1")
	control = UserController{ApplicationController{}}
}

func tearDownSignup() {
	c = &gin.Context{}
}

func TestSignup(t *testing.T) {
	setupSignup()

	defer tearDownSignup()
	control.Signup(c)
	if w.Code == 200 {
		t.Log("signup is ok")
	} else {
		t.Log("signup is error")
	}

}
