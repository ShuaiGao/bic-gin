package service

import (
	"bic-gin/internal/schema"
	"bic-gin/pkg/gen"
	"bic-gin/pkg/gen/api"
	"bic-gin/pkg/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AuthService struct {
}

func checkAuth(username, password string) (schema.User, bool, error) {
	return schema.User{
		Model:    gorm.Model{ID: 10},
		Username: username,
	}, true, nil
}

func (as AuthService) PostAuth(ctx *gin.Context, in *api.RequestAuth) (out *api.ResponseAuth, code api.ErrCode) {
	code = api.ECSuccess
	if err := gen.Validate.Struct(in); err != nil {
		code = api.ECParam
		return
	}
	user, ok, err := checkAuth(in.Username, in.Password)
	if !ok || err != nil {
		code = api.ECAuth.Wrap(err)
		return
	}
	token, tokenRefresh, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		code = api.ECTokenGen.Wrap(err)
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
