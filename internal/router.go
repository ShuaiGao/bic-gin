package internal

import (
	_ "bic-gin/docs"
	"bic-gin/internal/service"
	"bic-gin/pkg/gen/api"
	"bic-gin/pkg/jwt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	if gin.IsDebugging() {
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	apiNoJwt := g.Group("/api/bic-gin")
	api.RegisterAuthServiceHttpHandler(apiNoJwt, service.AuthService{})
	apiJWT := g.Group("/api/bic-gin", jwt.JWT())

	//注册、设置路由
	api.RegisterUserServiceHttpHandler(apiJWT, service.UserService{})
	api.RegisterAdminServiceHttpHandler(apiJWT, service.AdminService{})
	return g
}
