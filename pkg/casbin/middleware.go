package guarder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var prefix string
var prefixLen int

func SetPrefix(urlPrefix string) {
	prefix = urlPrefix
	prefixLen = len(urlPrefix)
}

// RequireURLPermission url 权限控制中间件
func RequireURLPermission(c *gin.Context) {
	obj := c.FullPath()
	if len(obj) < prefixLen {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	obj = obj[prefixLen:]
	sub := c.MustGet(UserSub)
	action := c.Request.Method
	ok, err := Enforcer().Enforce(sub, obj, action)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}

// RequireObjPermission restful 风格的url obj权限控制中间件
// TODO 该中间件暂不支持
func RequireObjPermission(objType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.MustGet(UserSub)
		method := c.Request.Method
		id := c.Param(ID)
		if id == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		obj := fmt.Sprintf("%s:%s", objType, method)
		ok, err := Enforcer().Enforce(sub, obj, id)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
