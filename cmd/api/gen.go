package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Swagger struct {
	Paths map[string]SwaggerItem `json:"paths"`
}
type SwaggerItem struct {
	Post   *SwaggerMethodItem `json:"post"`
	Get    *SwaggerMethodItem `json:"get"`
	Patch  *SwaggerMethodItem `json:"patch"`
	Delete *SwaggerMethodItem `json:"delete"`
}

func (si *SwaggerItem) GetMethods() []*SwaggerMethodItem {
	var ret []*SwaggerMethodItem
	if si.Post != nil {
		si.Post.method = "POST"
		ret = append(ret, si.Post)
	}
	if si.Get != nil {
		si.Get.method = "GET"
		ret = append(ret, si.Get)
	}
	if si.Patch != nil {
		si.Patch.method = "PATCH"
		ret = append(ret, si.Patch)
	}
	if si.Delete != nil {
		si.Delete.method = "DELETE"
		ret = append(ret, si.Delete)
	}
	return ret
}

type SwaggerMethodItem struct {
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
	method  string
}

func DealApiData() {
	data, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		panic(err)
	}
	swaggerInfo := &Swagger{}
	err = json.Unmarshal(data, swaggerInfo)
	if err != nil {
		panic(err)
	}

	var builder []string
	for k, v := range swaggerInfo.Paths {
		methods := v.GetMethods()
		if len(methods) == 0 {
			fmt.Println("not found method ", k)
			return
		}
		for _, v := range methods {
			key := fmt.Sprintf("%s-%s", k, v.method)
			builder = append(builder,
				"\t"+fmt.Sprintf(`{Key: "%s", Method: "%s", Url: "%s", Label: "%s"},`,
					key, v.method, k, v.Summary))
		}
	}
	sort.Strings(builder)
	var permissionStr = `package data

import "bic-gin/internal/schema"

var initApiData = []*schema.Api{
` + strings.Join(builder, "\n") + `
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
`
	err = os.WriteFile("internal/schema/data/api.go", []byte(permissionStr), 0666)
	if err != nil {
		fmt.Println("write file error ", err)
	}
}

func main() {
	DealApiData()
}
