// Code generated by protoc-gen-bic. DO NOT EDIT.
// versions:2.4.2

package api

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthServiceHttpHandler(g *gin.RouterGroup, srvs AuthService) {
	tmp := &x_AuthService{xx: srvs}
	g.POST("/v1/auth/", tmp.PostAuth)
	g.POST("/v1/token/", tmp.PostRefreshToken)
}

type AuthService interface {
	PostAuth(ctx *gin.Context, in *RequestAuth) (out *ResponseAuth, code ErrCode)
	PostRefreshToken(ctx *gin.Context, in *RequestRefreshToken) (out *ResponseAuth, code ErrCode)
}

// generated http handle
type AuthServiceHttpHandler interface {
	PostAuth(ctx *gin.Context)
	PostRefreshToken(ctx *gin.Context)
}

type x_AuthService struct {
	xx AuthService
}

// @Summary 获取用户列表
// @Tags    Auth-Service
// @Produce json
// @Param   username body     string true "用户名"
// @Param   password body     string true "密码"
// @Success 200      {object} ResponseAuth
// @Failure 401      {string} string "header need Authorization data"
// @Failure 403      {string} string "no api permission or no obj permission"
// @Router  /v1/auth/ [POST]
func (x *x_AuthService) PostAuth(ctx *gin.Context) {
	req := &RequestAuth{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "detail": "request error"})
		return
	}
	rsp, errCode := x.xx.PostAuth(ctx, req)

	ctx.JSON(http.StatusOK, gin.H{
		"code":   errCode.Code(),
		"detail": errCode.String(),
		"data":   rsp,
	})
}

// @Summary 刷新token
// @Tags    Auth-Service
// @Produce json
// @Param   token_refresh body     string true "参数无注释"
// @Success 200           {object} ResponseAuth
// @Failure 401           {string} string "header need Authorization data"
// @Failure 403           {string} string "no api permission or no obj permission"
// @Router  /v1/token/ [POST]
func (x *x_AuthService) PostRefreshToken(ctx *gin.Context) {
	req := &RequestRefreshToken{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "detail": "request error"})
		return
	}
	rsp, errCode := x.xx.PostRefreshToken(ctx, req)

	ctx.JSON(http.StatusOK, gin.H{
		"code":   errCode.Code(),
		"detail": errCode.String(),
		"data":   rsp,
	})
}