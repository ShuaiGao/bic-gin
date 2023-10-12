package service

import (
	"bic-gin/internal/schema"
	"bic-gin/pkg/gen"
	"bic-gin/pkg/gen/api"
	"bic-gin/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthService struct {
}

func checkAuth(username, password string) (schema.User, bool, error) {
	return schema.User{Username: username}, true, nil
}

func (as AuthService) PostAuth(ctx *gin.Context, in *api.RequestAuth) (out *api.ResponseAuth, code api.ErrCode) {
	code = api.ECSuccess
	if err := gen.Validate.Struct(in); err != nil {
		code = api.ECParam
		return
	}

	req := &api.RequestAuth{}
	if ok := ctx.Bind(req); ok != nil {
		ctx.JSON(http.StatusOK, gen.Response{Code: 400, Message: " request error"})
		return
	}
	if err := gen.Validate.Struct(req); err != nil {
		ctx.JSON(http.StatusOK, gen.Response{Code: 400, Message: "参数错误"})
		return
	}
	user, ok, err := checkAuth(req.Username, req.Password)
	if !ok || err != nil {
		ctx.JSON(http.StatusOK, gen.Response{Code: 400, Message: "username or password wrong"})
		return
	}
	token, tokenRefresh, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.JSON(http.StatusOK, gen.Response{Code: 500, Message: "gen token error"})
		return
	}
	out = &api.ResponseAuth{Token: token, TokenRefresh: tokenRefresh}
	return
}

func (as AuthService) PostRefreshToken(ctx *gin.Context, in *api.RequestRefreshToken) (out *api.ResponseAuth, code api.ErrCode) {
	code = api.ECSuccess
	claims, err := jwt.ParseRefreshToken(in.TokenRefresh)
	if err != nil {
		code = api.ECParamRefreshToken
		return
	}
	token, _, err := jwt.GenerateToken(claims.UserID, claims.Username)
	if err != nil {
		ctx.JSON(http.StatusOK, gen.Response{Code: 500, Message: "gen token error"})
		return
	}
	out = &api.ResponseAuth{Token: token, TokenRefresh: in.TokenRefresh}
	return
}
