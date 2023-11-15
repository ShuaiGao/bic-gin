package service

import (
	"bic-gin/internal/schema"
	"bic-gin/pkg/db"
	"bic-gin/pkg/gen/api"
	"bic-gin/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AdminService struct {
}

func (a AdminService) GetApis(ctx *gin.Context) (out *api.ResponseApis, code api.ErrCode) {
	code = api.ECSuccess
	var apis []*schema.Api
	if err := db.SqlDB().Find(&apis).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	out = &api.ResponseApis{}
	for _, v := range apis {
		out.ApiList = append(out.ApiList, &api.ApiItem{
			Key:    v.Key,
			Url:    v.Url,
			Method: v.Method,
			Label:  v.Label,
		})
	}
	return
}

// GetRoutes
// get route tree, there are two choose
// 1: based on user all permission, we use this way for example
// 2: based on user selected role
func (a AdminService) GetRoutes(ctx *gin.Context) (out *api.ResponseRoute, code api.ErrCode) {
	code = api.ECSuccess
	userId := ctx.GetUint(jwt.UserId)
	var roles []*schema.Role
	if err := db.SqlDB().Model(&schema.Role{}).
		Preload("MenuActions").
		Where("id in (select role_id from role_user where user_id = ?)", userId).
		Find(&roles, userId).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	var menus []*schema.Menu
	if err := db.SqlDB().Model(&schema.Menu{}).Find(&menus).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	menuMap := make(map[string][]*schema.Menu)
	for _, v := range menus {
		if _, ok := menuMap[v.FatherKey]; ok {
			menuMap[v.FatherKey] = append(menuMap[v.FatherKey], v)
		} else {
			menuMap[v.FatherKey] = []*schema.Menu{v}
		}
	}
	actionMap := make(map[string][]*schema.MenuAction)
	for _, v := range roles {
		for _, p := range v.MenuActions {
			if _, ok := actionMap[p.MenuKey]; ok {
				actionMap[p.MenuKey] = append(actionMap[p.MenuKey], p)
			} else {
				actionMap[p.MenuKey] = []*schema.MenuAction{p}
			}
		}
	}

	out = &api.ResponseRoute{
		ItemList: fillChildrenRoute(menuMap, actionMap, nil, ""),
	}
	return
}

func fillChildrenRoute(menuMap map[string][]*schema.Menu, actionMap map[string][]*schema.MenuAction, father *api.RouteItem, id string) []*api.RouteItem {
	children, ok := menuMap[id]
	if !ok {
		if father == nil {
			return nil
		}
		if actions, find := actionMap[id]; find {
			for _, v := range actions {
				father.Actions = append(father.Actions, v.Key)
			}
		}
		return nil
	}
	var itemList []*api.RouteItem
	for _, v := range children {
		item := &api.RouteItem{
			Key:  v.Key,
			Name: v.Name,
			Path: v.Path,
			Meta: &api.RouteMeta{
				Title: v.Label,
			},
		}
		item.Children = fillChildrenRoute(menuMap, actionMap, item, v.Key)
		itemList = append(itemList, item)
	}
	return itemList
}

func (a AdminService) GetRoles(ctx *gin.Context) (out *api.ResponseRoles, code api.ErrCode) {
	code = api.ECSuccess
	var dataList []*schema.Role
	if err := db.SqlDB().Find(&dataList).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	out = &api.ResponseRoles{}
	for _, v := range dataList {
		out.DataList = append(out.DataList, &api.RoleItem{
			Id:   uint32(v.ID),
			Name: v.Name,
		})
	}
	return
}

func (a AdminService) GetRole(ctx *gin.Context, id uint) (out *api.ResponseGetRole, code api.ErrCode) {
	code = api.ECSuccess
	var role schema.Role
	if err := db.SqlDB().Model(&role).
		Preload("MenuActions").
		Preload("MenuActions.Menu").
		First(&role, id).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	out = &api.ResponseGetRole{
		Id:   uint32(role.ID),
		Name: role.Name,
	}
	return
}

func (a AdminService) GetMenus(ctx *gin.Context) (out *api.ResponseGetMenus, code api.ErrCode) {
	code = api.ECSuccess
	var dataList []*schema.Menu
	if err := db.SqlDB().Find(&dataList).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	out = &api.ResponseGetMenus{}
	for _, v := range dataList {
		out.RouteList = append(out.RouteList, &api.ItemMenu{
			Id:    v.Key,
			Label: v.Label,
		})
	}
	return
}

func (a AdminService) PostMenuAction(ctx *gin.Context, in *api.RequestPostMenuAction, key string) (out *api.Empty, code api.ErrCode) {
	code = api.ECSuccess
	var menu schema.Menu
	if err := db.SqlDB().Model(&schema.Menu{}).Where("`key` = ?", key).First(&menu).Error; err != nil {
		code = api.ECDbFirst.Wrap(err)
		return
	}
	var count int64
	if err := db.SqlDB().Model(&schema.MenuAction{}).Where("`key` = ?", in.Key).Count(&count).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	action := &schema.MenuAction{
		Key:     in.Key,
		Label:   in.Label,
		MenuKey: key,
	}
	if len(in.Apis) > 0 {
		if err := db.SqlDB().Model(&schema.Api{}).Where("`key` in ?", in.Apis).Find(&action.Apis).Error; err != nil {
			code = api.ECDbFind.Wrap(err)
			return
		}
		if len(in.Apis) != len(action.Apis) {
			code = api.ECParam.Wrap("部分api key不存在")
			return
		}
	}
	if err := db.SqlDB().Create(action).Error; err != nil {
		code = api.ECDbCreate.Wrap(err)
		return
	}
	return
}

func (a AdminService) GetMenuAction(ctx *gin.Context, key string) (out *api.ResponseGetMenuAction, code api.ErrCode) {
	code = api.ECSuccess
	var actions []schema.MenuAction
	if err := db.SqlDB().
		Preload("Apis").
		Where("`menu_key` = ?", key).Find(&actions).Error; err != nil {
		code = api.ECDbFirst.Wrap(err)
		return
	}
	out = &api.ResponseGetMenuAction{}
	for _, v := range actions {
		item := &api.ItemMenuAction{
			Key:   v.Key,
			Label: v.Label,
		}
		for _, vv := range v.Apis {
			item.ApiList = append(item.ApiList, &api.ItemApi{
				Key:    vv.Key,
				Label:  vv.Label,
				Url:    vv.Url,
				Method: vv.Method,
			})
		}
		out.ActionList = append(out.ActionList, item)
	}

	return
}

func (a AdminService) PatchMenuAction(ctx *gin.Context, in *api.RequestPostMenuAction) (out *api.Empty, code api.ErrCode) {
	code = api.ECSuccess
	var action schema.MenuAction
	if err := db.SqlDB().Where("`key` = ?", in.Key).First(&action).Error; err != nil {
		code = api.ECDbFirst.Wrap(err)
		return
	}
	sql := "insert into menu_action_api (menu_action_key, api_key) values"
	for _, v := range in.Apis {
		sql += fmt.Sprintf("(%s,%s)", in.Key, v)
	}
	sql += ";"
	tx := db.SqlDB().Begin()
	if err := tx.Exec("delete from menu_action_api where menu_action_key = ?", in.Key).Error; err != nil {
		code = api.ECDbExec.Wrap(err)
		tx.Rollback()
		return
	}
	if err := tx.Exec(sql).Error; err != nil {
		code = api.ECDbExec.Wrap(err)
		tx.Rollback()
		return
	}
	// TODO 修改casbin中的权限项
	if err := tx.Commit().Error; err != nil {
		code = api.ECDbCommit.Wrap(err)
		return
	}
	return
}
