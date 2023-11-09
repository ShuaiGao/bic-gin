package data

import (
	"bic-gin/internal/schema"
	"bic-gin/pkg/db"
	"gorm.io/gorm/clause"
)

func InitData() {
	InitRole()
	InitUser()
	InitApi()
}

func InitApi() {
	_ = db.SqlDB().Clauses(clause.OnConflict{DoNothing: true}).Save(&initApiData).Error
}

func InitUser() {
	var count int64
	if err := db.SqlDB().Model(&schema.User{}).Count(&count).Error; err != nil || count > 0 {
		return
	}
	user := schema.User{
		Username: "bic",
		Name:     "BIC",
	}
	if err := db.SqlDB().Create(&user).Error; err != nil {
		return
	}
	var role schema.Role
	if err := db.SqlDB().Where("name = '管理员'").First(&role).Error; err != nil {
		return
	}
	db.SqlDB().Exec("insert into role_user(role_id, user_id) values (?,?)", role.ID, user.ID)
}

func InitRole() {
	var count int64
	if err := db.SqlDB().Model(&schema.Menu{}).Count(&count).Error; err != nil || count > 0 {
		return
	}
	var apis []*schema.Api
	apis = append(apis, []*schema.Api{
		{Url: "/api/bic-gin/v1/routes", Method: "GET"},
		{Url: "/api/bic-gin/v1/roles", Method: "GET"},
		{Url: "/api/bic-gin/v1/role/:id", Method: "GET"},
		{Url: "/api/bic-gin/v1/menus", Method: "GET"},
		{Url: "/api/bic-gin/v1/menu/:id/action", Method: "POST"},
	}...)
	for _, v := range apis {
		v.Key = v.Url + "-" + v.Method
	}
	var menus []*schema.Menu
	permission := &schema.Menu{
		Key:   "admin",
		Label: "权限管理",
		Rank:  10,
	}
	if err := db.SqlDB().Omit("FatherKey").Create(permission).Error; err != nil {
		return
	}
	adminUser := &schema.Menu{
		Key:       "admin-user",
		Label:     "用户管理",
		FatherKey: "admin",
		Rank:      1,
	}
	adminRole := &schema.Menu{
		Key:       "admin-role",
		Label:     "角色管理",
		FatherKey: "admin",
		Rank:      2,
	}
	adminMenu := &schema.Menu{
		Key:       "admin-menu",
		Label:     "菜单管理",
		FatherKey: "admin",
		Rank:      3,
	}
	menus = append(menus, adminUser, adminRole, adminMenu)
	if err := db.SqlDB().Create(menus).Error; err != nil {
		return
	}
	var menuActions []*schema.MenuAction
	menuActions = append(menuActions, []*schema.MenuAction{
		{
			Key:     "admin-user-view",
			Label:   "查询",
			MenuKey: adminUser.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/users", Method: "GET", Key: "/api/bic-gin/v1/users-GET"},
			},
		},
		{
			Key:     "admin-user-change",
			Label:   "修改",
			MenuKey: adminUser.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/user/:id", Method: "PATCH", Key: "/api/bic-gin/v1/user/:id-PATCH"},
			},
		},
		{
			Key:     "admin-user-ban",
			Label:   "禁用",
			MenuKey: adminUser.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/user/:id/ban", Method: "GET", Key: "/api/bic-gin/v1/user/:id/ban-GET"},
			},
		},
	}...)
	menuActions = append(menuActions, []*schema.MenuAction{
		{
			Key:     "admin-role-view",
			Label:   "查询",
			MenuKey: adminRole.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/roles", Method: "GET", Key: "/api/bic-gin/v1/roles-GET"},
			},
		},
		{
			Key:     "admin-role-create",
			Label:   "创建",
			MenuKey: adminRole.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/roles", Method: "POST", Key: "/api/bic-gin/v1/roles-POST"},
			},
		},
		{
			Key:     "admin-role-change",
			Label:   "修改",
			MenuKey: adminRole.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/role/:id", Method: "Patch", Key: "/api/bic-gin/v1/role/:id-Patch"},
			},
		},
	}...)
	menuActions = append(menuActions, []*schema.MenuAction{
		{
			Key:     "admin-menu-view",
			Label:   "查询",
			MenuKey: adminMenu.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/menus", Method: "GET", Key: "/api/bic-gin/v1/menus-GET"},
			},
		},
		{
			Key:     "admin-menu-create",
			Label:   "创建",
			MenuKey: adminMenu.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/menus", Method: "POST", Key: "/api/bic-gin/v1/menus-POST"},
			},
		},
		{
			Key:     "admin-menu-change",
			Label:   "修改",
			MenuKey: adminMenu.Key,
			Apis: []schema.Api{
				{Url: "/api/bic-gin/v1/menu/:id", Method: "Patch", Key: "/api/bic-gin/v1/menu/:id-Patch"},
			},
		},
	}...)
	if err := db.SqlDB().Create(&menuActions).Error; err != nil {
		return
	}
	admin := schema.Role{
		Name:        "管理员",
		MenuActions: menuActions,
	}
	if err := db.SqlDB().Create(&admin).Error; err != nil {
		return
	}
}
