package service

import (
	"bic-gin/internal/schema"
	"bic-gin/pkg/db"
	"bic-gin/pkg/gen/api"
	"bic-gin/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type AdminService struct {
}

// GetRoutes
// get route tree there are two choose
// 1: based on user all permission, we use this way for example
// 2: based on user selected role
func (a AdminService) GetRoutes(ctx *gin.Context) (out *api.ResponseRoute, code api.ErrCode) {
	userId := ctx.GetUint(jwt.UserId)
	var roles []*schema.Role
	if err := db.SqlDB().Model(&schema.Role{}).
		Preload("Roles.PageActions").
		Preload("Roles.PageActions.Menu").
		Where("id in (select role_id from role_user where user_id = ?)", userId).
		Find(&roles, userId).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	actionMap := make(map[string][]*schema.PageAction)
	menuMap := make(map[string][]*schema.Menu)
	for _, v := range roles {
		for _, p := range v.PageActions {
			if _, ok := actionMap[p.MenuKey]; ok {
				actionMap[p.MenuKey] = append(actionMap[p.MenuKey], &p)
			} else {
				actionMap[p.MenuKey] = []*schema.PageAction{&p}
			}

			if _, ok := menuMap[p.Menu.FatherKey]; ok {
				menuMap[p.Menu.FatherKey] = append(menuMap[p.Menu.FatherKey], &p.Menu)
			} else {
				menuMap[p.Menu.FatherKey] = []*schema.Menu{&p.Menu}
			}
		}
	}

	out = &api.ResponseRoute{
		ItemList: fillChildrenRoute(menuMap, actionMap, nil, ""),
	}
	return
}

func fillChildrenRoute(menuMap map[string][]*schema.Menu, actionMap map[string][]*schema.PageAction, father *api.RouteItem, id string) []*api.RouteItem {
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
			Name: v.Label,
		}
		fillChildrenRoute(menuMap, actionMap, item, v.Key)
		itemList = append(itemList, item)
	}
	return itemList
}

func (a AdminService) GetRoles(ctx *gin.Context) (out *api.ResponseRoles, code api.ErrCode) {
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
	var role schema.Role
	if err := db.SqlDB().Model(&role).
		Preload("PageActions").
		Preload("PageActions.Menu").
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
	var dataList []*schema.Menu
	if err := db.SqlDB().Find(&dataList).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	out = &api.ResponseGetMenus{}
	for _, v := range dataList {
		out.RouteList = append(out.RouteList, &api.MenuItem{
			Id:    v.Key,
			Label: v.Label,
		})
	}
	return
}

func (a AdminService) PostMenuPageAction(ctx *gin.Context, in *api.RequestPostMenuPageAction, key string) (out *api.Empty, code api.ErrCode) {
	var menu schema.Menu
	if err := db.SqlDB().Model(&schema.Menu{}).Where("`key` = ?", key).First(&menu).Error; err != nil {
		code = api.ECDbFirst.Wrap(err)
		return
	}
	var count int64
	if err := db.SqlDB().Model(&schema.PageAction{}).Where("`key` = ?", in.Key).Count(&count).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	action := &schema.PageAction{
		Key:     in.Key,
		Label:   in.Label,
		MenuKey: key,
	}
	for _, v := range in.Methods {
		action.Apis = append(action.Apis, schema.Api{
			Label: v.Note,
		})
	}
	if err := db.SqlDB().Create(action).Error; err != nil {
		code = api.ECDbCreate.Wrap(err)
		return
	}
	return
}
