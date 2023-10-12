package internal

import (
	"gin-bic/internal/service"
	"gin-bic/pkg/gen/api"
	"gin-bic/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}

// SetupRouter 启动路由，添加中间件
func SetupRouter() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(Cors())
	g.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})

	apiNoJwt := g.Group("/api/gin-bic")
	api.RegisterAuthServiceHttpHandler(apiNoJwt, service.AuthService{})
	apiJWT := g.Group("/api/gin-bic", jwt.JWT())

	//注册、设置路由
	api.RegisterUserServiceHttpHandler(apiJWT, service.UserService{})
	return g
}
