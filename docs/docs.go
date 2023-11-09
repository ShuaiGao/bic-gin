// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "GaoZiJia",
            "email": "boringmanman@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/apis": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "获取接口列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseApis"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/auth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth-Service"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "用户名",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "密码",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseAuth"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/menu/:key/action": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "获取菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some id",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseGetMenuAction"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "添加菜单页面行为",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some id",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "参数无注释",
                        "name": "key",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "参数无注释",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "对应api key列表",
                        "name": "apis",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/menu/action": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "修改菜单页面行为",
                "parameters": [
                    {
                        "description": "参数无注释",
                        "name": "key",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "参数无注释",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "对应api key列表",
                        "name": "apis",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/menus": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "获取菜单列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseGetMenus"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/role/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "获取角色详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "some id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseGetRole"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/roles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "获取角色列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseRoles"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/routes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Service"
                ],
                "summary": "获取路由",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseRoute"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/token/refresh": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth-Service"
                ],
                "summary": "刷新token",
                "parameters": [
                    {
                        "description": "参数无注释",
                        "name": "token_refresh",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseAuth"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/:id": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Service"
                ],
                "summary": "修改用户权限",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "some id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "参数无注释",
                        "name": "role_id",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Service"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "参数无注释",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "参数无注释",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ResponseUsers"
                        }
                    },
                    "401": {
                        "description": "header need Authorization data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "no api permission or no obj permission",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ApiItem": {
            "type": "object",
            "properties": {
                "key": {
                    "description": "required",
                    "type": "string"
                },
                "label": {
                    "description": "required",
                    "type": "string"
                },
                "method": {
                    "description": "required",
                    "type": "string"
                },
                "url": {
                    "description": "required",
                    "type": "string"
                }
            }
        },
        "api.HttpMethod": {
            "type": "object",
            "properties": {
                "method": {
                    "description": "required",
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "url": {
                    "description": "required",
                    "type": "string"
                }
            }
        },
        "api.ItemApi": {
            "type": "object",
            "properties": {
                "Key": {
                    "description": "required",
                    "type": "string"
                },
                "Label": {
                    "description": "required",
                    "type": "string"
                },
                "Method": {
                    "description": "required",
                    "type": "string"
                },
                "Url": {
                    "description": "required",
                    "type": "string"
                }
            }
        },
        "api.ItemMenu": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ItemMenu"
                    }
                },
                "id": {
                    "description": "required",
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.MenuPageAction"
                    }
                },
                "label": {
                    "description": "required",
                    "type": "string"
                },
                "menu_item_type": {
                    "description": "required",
                    "type": "integer"
                }
            }
        },
        "api.ItemMenuAction": {
            "type": "object",
            "properties": {
                "Api_list": {
                    "description": "对应http接口",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ItemApi"
                    }
                },
                "key": {
                    "description": "required",
                    "type": "string"
                },
                "label": {
                    "description": "required",
                    "type": "string"
                }
            }
        },
        "api.MenuPageAction": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "required",
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "methods": {
                    "description": "对应http接口",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.HttpMethod"
                    }
                }
            }
        },
        "api.PageItem": {
            "type": "object",
            "properties": {
                "checked": {
                    "type": "boolean"
                },
                "id": {
                    "description": "required",
                    "type": "string"
                },
                "label": {
                    "type": "string"
                }
            }
        },
        "api.ResponseApis": {
            "type": "object",
            "properties": {
                "api_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ApiItem"
                    }
                }
            }
        },
        "api.ResponseAuth": {
            "type": "object",
            "required": [
                "expires",
                "expires_refresh",
                "token",
                "token_refresh"
            ],
            "properties": {
                "expires": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "integer"
                },
                "expires_refresh": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "integer"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "token": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "string"
                },
                "token_refresh": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.ResponseGetMenuAction": {
            "type": "object",
            "properties": {
                "action_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ItemMenuAction"
                    }
                }
            }
        },
        "api.ResponseGetMenus": {
            "type": "object",
            "properties": {
                "route_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ItemMenu"
                    }
                }
            }
        },
        "api.ResponseGetRole": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "create_time": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "create_user": {
                    "description": "创建用户",
                    "type": "string"
                },
                "father_id": {
                    "description": "父角色ID",
                    "type": "integer"
                },
                "father_name": {
                    "description": "父角色名",
                    "type": "string"
                },
                "id": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "integer"
                },
                "name": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "route_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.RouteNode"
                    }
                },
                "update_time": {
                    "description": "更新时间",
                    "type": "integer"
                },
                "update_user": {
                    "description": "更新用户",
                    "type": "string"
                }
            }
        },
        "api.ResponseRoles": {
            "type": "object",
            "properties": {
                "data_list": {
                    "description": "uint32 page = 1;\nuint32 page_size = 2;\nuint32 total = 3;",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.RoleItem"
                    }
                }
            }
        },
        "api.ResponseRoute": {
            "type": "object",
            "properties": {
                "item_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.RouteItem"
                    }
                }
            }
        },
        "api.ResponseUsers": {
            "type": "object",
            "properties": {
                "data_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.User"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "api.RoleItem": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "create_time": {
                    "description": "创建时间",
                    "type": "integer"
                },
                "create_user": {
                    "description": "创建用户",
                    "type": "string"
                },
                "father_id": {
                    "description": "父角色ID",
                    "type": "integer"
                },
                "father_name": {
                    "description": "父角色名",
                    "type": "string"
                },
                "id": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "integer"
                },
                "name": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "update_time": {
                    "description": "更新时间",
                    "type": "integer"
                },
                "update_user": {
                    "description": "更新用户",
                    "type": "string"
                }
            }
        },
        "api.RouteItem": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.RouteItem"
                    }
                },
                "key": {
                    "type": "string"
                },
                "meta": {
                    "$ref": "#/definitions/api.RouteMeta"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "api.RouteMeta": {
            "type": "object",
            "properties": {
                "auths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "icon": {
                    "type": "string"
                },
                "rank": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.RouteNode": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.RouteNode"
                    }
                },
                "id": {
                    "description": "required",
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.PageItem"
                    }
                },
                "label": {
                    "type": "string"
                }
            }
        },
        "api.User": {
            "type": "object",
            "properties": {
                "ban": {
                    "description": "账号是否禁用",
                    "type": "boolean"
                },
                "email": {
                    "description": "required",
                    "type": "string"
                },
                "id": {
                    "description": "required",
                    "type": "integer"
                },
                "name": {
                    "description": "姓名",
                    "type": "string"
                },
                "update_time": {
                    "description": "required",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.bic-gin.cn",
	BasePath:         "/api/bic-gin",
	Schemes:          []string{},
	Title:            "bic-gin api文档",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
