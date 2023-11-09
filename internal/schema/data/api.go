package data

import "bic-gin/internal/schema"

var initApiData = []*schema.Api{
	{Key: "/v1/auth-POST", Method: "POST", Url: "/v1/auth", Label: "登录"},
	{Key: "/v1/menu/:key/action-POST", Method: "POST", Url: "/v1/menu/:key/action", Label: "添加菜单页面行为"},
	{Key: "/v1/menus-GET", Method: "GET", Url: "/v1/menus", Label: "获取菜单列表"},
	{Key: "/v1/role/:id-GET", Method: "GET", Url: "/v1/role/:id", Label: "获取角色详情"},
	{Key: "/v1/roles-GET", Method: "GET", Url: "/v1/roles", Label: "获取角色列表"},
	{Key: "/v1/routes-GET", Method: "GET", Url: "/v1/routes", Label: "获取路由"},
	{Key: "/v1/token/refresh-POST", Method: "POST", Url: "/v1/token/refresh", Label: "刷新token"},
	{Key: "/v1/user/:id-PATCH", Method: "PATCH", Url: "/v1/user/:id", Label: "修改用户权限"},
	{Key: "/v1/users-GET", Method: "GET", Url: "/v1/users", Label: "获取用户列表"},
}

var ApiMap map[string]*schema.Api

func GetApi(key string) *schema.Api {
	return ApiMap[key]
}

func init() {
	ApiMap = make(map[string]*schema.Api)
	for _, v := range initApiData {
		ApiMap[v.Key] = v
	}
}
