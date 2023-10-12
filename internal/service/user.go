package service

import (
	"gin-bic/pkg/gen/api"
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (UserService) GetUsers(ctx *gin.Context, in *api.RequestUsers) (out *api.ResponseUsers, code api.ErrCode) {
	code = api.ECSuccess
	return
}
