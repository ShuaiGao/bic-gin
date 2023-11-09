package guarder

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ResourceType uint
type ShowType uint

var (
	ShowTypeSelect   ShowType = 1 // 下拉列表
	ShowTypeCascader ShowType = 2 // 级联选择
)

type ResourceItem struct {
	Type          ResourceType
	ShowType      ShowType
	Resource      ResourceTable
	CasbinObjName CasbinObjName
}

var ResourceMap = make(map[ResourceType]ResourceItem)
var ResourceKeyList []ResourceType

// RegisterResourceTable 注册资源权限，注册顺序即前端显示顺序
func RegisterResourceTable(item ...ResourceItem) {
	for _, v := range item {
		ResourceKeyList = append(ResourceKeyList, v.Type)
		ResourceMap[v.Type] = v
	}
}

type ResourceTreeItem struct {
	Id       uint32              `json:"id"`
	Label    string              `json:"label"`
	Children []*ResourceTreeItem `json:"children"`
}

type ResourceTable interface {
	TableName() string
	GetResourceTree() ([]*ResourceTreeItem, map[uint]string, error)
}

type CasbinResource struct {
	gorm.Model
	CasbinKey string `gorm:"not null;uniqueIndex:idx_casbin_key;size:32"`
}

var errIdCheck = errors.New("permission check error")

func CheckAndCasbinKeys(db *gorm.DB, rType ResourceType, idList []uint32) ([]string, error) {
	tableName := GetTableName(rType)
	var resourceList []*CasbinResource
	if err := db.Table(tableName).
		Select("id", "casbin_key").Where("id in ?", idList).
		Find(&resourceList).Error; err != nil || len(resourceList) != len(idList) {
		return nil, errIdCheck
	}
	var ret []string
	for _, v := range resourceList {
		ret = append(ret, v.CasbinKey)
	}
	return ret, nil
}

func GetTableName(t ResourceType) string {
	if _, ok := ResourceMap[t]; !ok {
		return ""
	}
	item := ResourceMap[t]
	return item.Resource.TableName()
}

// SetMockResourcePermission 设置casbin mock
func SetMockResourcePermission[T string | uint | int](ctx *gin.Context, rt ResourceType, idList []T) error {
	userSub := ctx.GetString(UserSub)
	username := ctx.GetString(User)
	resourceGroupSub := GetGroupResourceSub(username)
	// 为用户添加资源权限组
	_, err := Enforcer().AddRoleForUser(userSub, resourceGroupSub)
	if err != nil {
		return err
	}
	resourceItem, ok := ResourceMap[rt]
	if !ok {
		return errors.New("not this ResourceType")
	}
	for _, v := range idList {
		action := fmt.Sprintf("%v", v)
		if _, err := Enforcer().AddPermissionsForUser(resourceGroupSub, []string{string(resourceItem.CasbinObjName), action}); err != nil {
			return err
		}
	}
	return nil
}
