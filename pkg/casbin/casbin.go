package guarder

import (
	"fmt"
	sa "github.com/ShuaiGao/string-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
)

const UserSub = "user_sub"
const User = "user"
const ID = "id"

type CasbinObjName string

var (
	CasbinGroupResourcePrefix = "g:r:"
	CasbinResourceAll         = "r:all"
	UserPrefix                = "u"
	GroupPrefix               = "g"
	redisWatcherAddress       string
	redisWatcherPassword      string
	redisChannelName          string
)

var enforcer *casbin.Enforcer

func Enforcer() *casbin.Enforcer {
	return enforcer
}

func SetRedisWatcherConfig(address, password, channelName string) {
	redisWatcherAddress = address
	redisWatcherPassword = password
	redisChannelName = channelName
}

func GetEffect(e string) string {
	if e == "deny" {
		return "deny"
	}
	return "allow"
}

func newModel(superUsername string) model.Model {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	if superUsername != "" {
		m.AddDef("m", "m", fmt.Sprintf("g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == \"u:%s\"", superUsername))
	} else {
		m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")
	}
	return m
}

func Setup(superUsername string, adapter persist.Adapter) {
	var err error
	enforcer, err = casbin.NewEnforcer(newModel(superUsername), adapter)
	if err != nil {
		panic("error: build casbin: " + err.Error())
	}
	if redisWatcherAddress != "" {
		w, _ := rediswatcher.NewWatcher(redisWatcherAddress, rediswatcher.WatcherOptions{
			Options: redis.Options{
				Network:  "tcp",
				Password: redisWatcherPassword,
			},
			Channel: redisChannelName,
			// Only exists in test, generally be true
			IgnoreSelf: true,
		})
		enforcer.SetWatcher(w)
		enforcer.EnableAutoNotifyWatcher(true)
		w.SetUpdateCallback(rediswatcher.DefaultUpdateCallback(enforcer))
	}
}

func SetupMock() {
	b := "" // b stores Casbin policy in JSON bytes.
	a := sa.NewAdapter(&b)
	Setup("", a)
}

func GetDenyColumn(user string, ModelName string) []string {
	var column []string
	policyList := Enforcer().GetFilteredPolicy(0, user, fmt.Sprintf("column:deny:%s", ModelName))
	for _, p := range policyList {
		column = append(column, p[2])
	}
	roles, err := Enforcer().GetImplicitRolesForUser(user)
	if err != nil {
		return []string{}
	}
	for _, role := range roles {
		policyList = Enforcer().GetFilteredPolicy(0, role, fmt.Sprintf("column:deny:%s", ModelName))
		for _, p := range policyList {
			column = append(column, p[2])
		}
	}
	return column
}

func GetCasbinResourceAll(name CasbinObjName) (obj string, action string) {
	return CasbinResourceAll, string(name)
}

func GetUserSub(username string) string {
	return fmt.Sprintf("%s:%s", UserPrefix, username)
}
func GetGroupSub(roleKey string) string {
	return fmt.Sprintf("%s:%s", GroupPrefix, roleKey)
}
func GetGroupResourceSub(key string) string {
	return fmt.Sprintf("%s%s", CasbinGroupResourcePrefix, key)
}

// GetAllowMapByUser 返回 有权限的数据id Map，是否有所有权限
func GetAllowMapByUser(userSub string, modelName CasbinObjName) (map[string]bool, bool) {
	obj, action := GetCasbinResourceAll(modelName)
	if ok, err := Enforcer().Enforce(userSub, obj, action); err == nil && ok {
		return nil, true
	}
	ret := make(map[string]bool)
	groupList, err := Enforcer().GetImplicitRolesForUser(userSub)
	if err != nil {
		return nil, false
	}
	for _, group := range groupList {
		itemList := Enforcer().GetFilteredPolicy(0, group, string(modelName))
		for _, v := range itemList {
			if len(v) < 3 {
				continue
			}
			ret[v[2]] = true
		}
	}
	return ret, false
}

// GetAllowListStringByUser 返回 有权限的数据id列表，是否有所有权限
func GetAllowListStringByUser(userSub string, modelName CasbinObjName) ([]string, bool) {
	obj, action := GetCasbinResourceAll(modelName)
	if ok, err := Enforcer().Enforce(userSub, obj, action); err == nil && ok {
		return nil, true
	}
	var ret []string
	groupList, err := Enforcer().GetImplicitRolesForUser(userSub)
	if err != nil {
		return nil, false
	}
	for _, group := range groupList {
		itemList := Enforcer().GetFilteredPolicy(0, group, string(modelName))
		for _, v := range itemList {
			if len(v) < 3 {
				continue
			}
			ret = append(ret, v[2])
		}
	}
	return ret, false
}

// GetAllowListStringByGroup 返回 有权限的数据id列表，是否有所有权限
func GetAllowListStringByGroup(groupSub string, modelName CasbinObjName) ([]string, bool) {
	obj, action := GetCasbinResourceAll(modelName)
	if ok, err := Enforcer().Enforce(groupSub, obj, action); err == nil && ok {
		return nil, true
	}
	var ret []string
	itemList := Enforcer().GetFilteredPolicy(0, groupSub, string(modelName))
	for _, v := range itemList {
		if len(v) < 3 {
			continue
		}
		ret = append(ret, v[2])
	}
	return ret, false
}

func GetAllowListUintByGroup(groupSub string, modelName CasbinObjName) ([]uint, bool) {
	obj, action := GetCasbinResourceAll(modelName)
	if ok, err := Enforcer().Enforce(groupSub, obj, action); err == nil && ok {
		return nil, true
	}
	var ret []uint
	itemList := Enforcer().GetFilteredPolicy(0, groupSub, string(modelName))
	for _, v := range itemList {
		if len(v) < 3 {
			continue
		}
		if n, err := strconv.Atoi(v[2]); err == nil {
			ret = append(ret, uint(n))
		}
	}
	return ret, false
}

func GetAllowListUintByUser(userSub string, modelName CasbinObjName) ([]uint, bool) {
	obj, action := GetCasbinResourceAll(modelName)
	if ok, err := Enforcer().Enforce(userSub, obj, action); err == nil && ok {
		return nil, true
	}
	var ret []uint
	groupList, err := Enforcer().GetImplicitRolesForUser(userSub)
	if err != nil {
		return nil, false
	}
	for _, group := range groupList {
		if !strings.HasPrefix(group, CasbinGroupResourcePrefix) {
			continue
		}
		itemList := Enforcer().GetFilteredPolicy(0, group, string(modelName))
		for _, v := range itemList {
			if len(v) < 3 {
				continue
			}
			if n, err := strconv.Atoi(v[2]); err == nil {
				ret = append(ret, uint(n))
			}
		}
	}
	return ret, false
}
