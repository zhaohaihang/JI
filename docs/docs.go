// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Zhao Haihang",
            "url": "http://www.swagger.io/support",
            "email": "1932859223@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/activity": {
            "post": {
                "description": "创建活动接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activity"
                ],
                "summary": "创建活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization header parameter",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "activity create info",
                        "name": "CreateActivityInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateActivityInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/serializer.Activity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/activity/near": {
            "get": {
                "description": "根据位置和半径查看活动接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activity"
                ],
                "summary": "查看指定点周围的所有活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "lat",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "lng",
                        "name": "lng",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "rad",
                        "name": "rad",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/serializer.DataList"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "item": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/serializer.Activity"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/activity/{aid}": {
            "get": {
                "description": "查看用户详情接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activity"
                ],
                "summary": "查看活动详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "activity ID",
                        "name": "aid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/serializer.Activity"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "put": {
                "description": "用户更新信息接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户更新信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization header parameter",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "user update info",
                        "name": "UpdateUserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/serializer.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/avatar": {
            "post": {
                "description": "上传用户头像接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "上传用户头像",
                "parameters": [
                    {
                        "type": "file",
                        "description": "图片文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization header parameter",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/serializer.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/login": {
            "post": {
                "description": "用户登录接口，如果用户不存在则创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "login user info",
                        "name": "LoginUserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.LoginUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/serializer.TokenData"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "user": {
                                                            "$ref": "#/definitions/serializer.User"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/{uid}": {
            "get": {
                "description": "查看用户信息接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查看用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization header parameter",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "user ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/serializer.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/{uid}/activity": {
            "get": {
                "description": "根据用户ID查看该用户创建的活动",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查看指定用户创建的所有活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page num",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "user ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/serializer.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/serializer.DataList"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "item": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/serializer.Activity"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Point": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                }
            }
        },
        "serializer.Activity": {
            "type": "object",
            "properties": {
                "current_Number": {
                    "type": "integer"
                },
                "endTime": {
                    "type": "integer"
                },
                "expected_number": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/serializer.Point"
                },
                "startTime": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user_avatar": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "serializer.CreateActivityInfo": {
            "type": "object",
            "required": [
                "end_time",
                "start_time",
                "status"
            ],
            "properties": {
                "end_time": {
                    "type": "integer"
                },
                "expected_number": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string",
                    "maxLength": 1000
                },
                "location": {
                    "$ref": "#/definitions/serializer.Point"
                },
                "start_time": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer",
                    "enum": [
                        1,
                        2,
                        3
                    ]
                },
                "title": {
                    "type": "string",
                    "maxLength": 30
                }
            }
        },
        "serializer.DataList": {
            "type": "object",
            "properties": {
                "item": {},
                "total": {
                    "type": "integer"
                }
            }
        },
        "serializer.LoginUserInfo": {
            "type": "object",
            "required": [
                "password",
                "type",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                },
                "type": {
                    "type": "integer",
                    "enum": [
                        1,
                        2
                    ]
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "serializer.Point": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                }
            }
        },
        "serializer.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "serializer.TokenData": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {}
            }
        },
        "serializer.UpdateUserInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "maxLength": 1000
                },
                "biography": {
                    "type": "string",
                    "maxLength": 1000
                },
                "email": {
                    "type": "string"
                },
                "extra": {
                    "type": "string",
                    "maxLength": 1000
                },
                "location": {
                    "$ref": "#/definitions/serializer.Point"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "serializer.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "biography": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "extra": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastLogin": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/model.Point"
                },
                "phone": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "JI API",
	Description:      "The api docs of JI project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
