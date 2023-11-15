// Code generated by protoc-gen-bic. DO NOT EDIT.
// versions:2.0.3

package api

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserServiceHttpHandler(g *gin.RouterGroup, srvs UserService) {
	tmp := &x_UserService{xx: srvs}
	g.GET("/v1/users", tmp.GetUsers)
	g.PATCH("/v1/user/:id", tmp.PatchUserPermission)
}

type UserService interface {
	GetUsers(ctx *gin.Context, in *RequestUsers) (out *ResponseUsers, code ErrCode)
	PatchUserPermission(ctx *gin.Context, in *RequestPatchPermission, id uint) (out *Empty, code ErrCode)
}

// generated http handle
type UserServiceHttpHandler interface {
	GetUsers(ctx *gin.Context)
	PatchUserPermission(ctx *gin.Context)
}

type x_UserService struct {
	xx UserService
}

// @Summary 获取用户列表
// @Tags    User-Service
// @Produce json
// @Param   page      query    uint32        true  "页码"
// @Param   page_size query    uint32        true  "每页数量"
// @Param   username  query    string        false "参数无注释"
// @Param   email     query    string        false "参数无注释"
// @Param   ban       query    UserBanStatus false "参数无注释"
// @Success 200       {object} ResponseUsers
// @Failure 401       {string} string "header need Authorization data"
// @Failure 403       {string} string "no api permission or no obj permission"
// @Router  /v1/users [GET]
func (x *x_UserService) GetUsers(ctx *gin.Context) {
	req := &RequestUsers{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "detail": "request error"})
		return
	}
	rsp, errCode := x.xx.GetUsers(ctx, req)

	ctx.JSON(http.StatusOK, gin.H{
		"code":   errCode.Code(),
		"detail": errCode.String(),
		"data":   rsp,
	})
}

// @Summary 修改用户权限
// @Tags    User-Service
// @Produce json
// @Param   id      path     uint   true  "some id"
// @Param   role_id body     uint32 false "参数无注释"
// @Success 200     {object} object null
// @Failure 401     {string} string "header need Authorization data"
// @Failure 403     {string} string "no api permission or no obj permission"
// @Router  /v1/user/:id [PATCH]
func (x *x_UserService) PatchUserPermission(ctx *gin.Context) {
	req := &RequestPatchPermission{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "detail": "request error"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   400,
			"detail": "param id should be int",
		})
	}
	rsp, errCode := x.xx.PatchUserPermission(ctx, req, uint(id))

	ctx.JSON(http.StatusOK, gin.H{
		"code":   errCode.Code(),
		"detail": errCode.String(),
		"data":   rsp,
	})
}
