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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/admin/login": {
            "post": {
                "description": "Admin Login by account and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin login",
                "parameters": [
                    {
                        "description": "Admin login request",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Admin login response",
                        "schema": {
                            "$ref": "#/definitions/response.AdminLoginResp"
                        }
                    }
                }
            }
        },
        "/company/advt/upload": {
            "put": {
                "description": "Company upload advertisement",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company upload advertisement",
                "responses": {
                    "200": {
                        "description": "Company upload advertisement request body",
                        "schema": {
                            "$ref": "#/definitions/request.CompanyUploadAdvtReq"
                        }
                    }
                }
            }
        },
        "/company/cancel": {
            "post": {
                "description": "Company Cancel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company Cancel",
                "responses": {
                    "200": {
                        "description": "nil"
                    }
                }
            }
        },
        "/company/get-info": {
            "get": {
                "description": "Company get information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company get information",
                "responses": {
                    "200": {
                        "description": "company info",
                        "schema": {
                            "$ref": "#/definitions/response.CompanyGetInfoResp"
                        }
                    }
                }
            }
        },
        "/company/login": {
            "post": {
                "description": "Company Login by account and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company login",
                "parameters": [
                    {
                        "description": "company login request",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CompanyLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Company login response",
                        "schema": {
                            "$ref": "#/definitions/response.CompanyLoginResp"
                        }
                    }
                }
            }
        },
        "/company/logout": {
            "post": {
                "description": "Company logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company logout",
                "responses": {
                    "200": {
                        "description": "response"
                    }
                }
            }
        },
        "/company/register": {
            "post": {
                "description": "Company register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company register",
                "parameters": [
                    {
                        "description": "company register request",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CompanyRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response"
                    }
                }
            }
        },
        "/company/search": {
            "get": {
                "description": "Get Companies by term",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Get Companies by term",
                "parameters": [
                    {
                        "type": "string",
                        "description": "term",
                        "name": "term",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Enum: 0 -\u003e Account, 1 -\u003e Name, 2 -\u003e Address, 3 -\u003e ManagerName, 4 -\u003e MangerTel, 5 -\u003e BusinessLicenseNumber",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Company Info",
                        "schema": {
                            "$ref": "#/definitions/response.GetCompaniesResp"
                        }
                    }
                }
            }
        },
        "/company/update-info": {
            "post": {
                "description": "Company update info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company update info",
                "parameters": [
                    {
                        "description": "company update info  request",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CompanyUpdateInfoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Company UpdateInfoReq response"
                    }
                }
            }
        },
        "/company/update-pwd": {
            "post": {
                "description": "Company update password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company update password",
                "parameters": [
                    {
                        "description": "company update password request",
                        "name": "request_body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.CompanyUpdatePwdReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Company update password response"
                    }
                }
            }
        },
        "/manage/admin/create": {
            "post": {
                "description": "Create Admin with account and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Admin create",
                "parameters": [
                    {
                        "description": "Admin create request",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Admin create response",
                        "schema": {
                            "$ref": "#/definitions/response.AdminCreateResp"
                        }
                    }
                }
            }
        },
        "/manage/advertisement/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Get all advertisements to be reviewed",
                "responses": {
                    "200": {
                        "description": "All advertisements",
                        "schema": {
                            "$ref": "#/definitions/response.GetAdvertisementsResp"
                        }
                    }
                }
            }
        },
        "/manage/company/info": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Allow Update For Companies",
                "parameters": [
                    {
                        "description": "allow companies update request",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AllowCompaniesUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "allow update for companies response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/manage/company/info_review_count": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Get info pending review companies count",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.GetCompaniesCountResp"
                        }
                    }
                }
            }
        },
        "/manage/company/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Get all companies",
                "responses": {
                    "200": {
                        "description": "Get companies response",
                        "schema": {
                            "$ref": "#/definitions/response.GetCompaniesResp"
                        }
                    }
                }
            }
        },
        "/manage/company/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Allow Registration For Companies",
                "parameters": [
                    {
                        "description": "allow companies register request",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AllowCompaniesRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "allow companies response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/manage/company/review": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Get all companies to be reviewed",
                "responses": {
                    "200": {
                        "description": "Get companies to be reviewed response",
                        "schema": {
                            "$ref": "#/definitions/response.GetCompaniesToBeReviewedResp"
                        }
                    }
                }
            }
        },
        "/manage/company/review_count": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "Get info pending review companies count",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.GetCompaniesCountResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "advertisement_model.AdvertisementInfo": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string",
                    "example": "2023-06-04"
                },
                "jump_to_url": {
                    "type": "string",
                    "example": "example.com/1919810.jpg"
                },
                "position": {
                    "type": "integer",
                    "example": 1
                },
                "start_date": {
                    "type": "string",
                    "example": "2023-05-14"
                },
                "title": {
                    "type": "string",
                    "example": "advertisement example"
                }
            }
        },
        "advertisement_model.AdvertisementToBePreviewedInfo": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "integer"
                },
                "end_date": {
                    "type": "string",
                    "example": "2023-06-04"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "jump_to_url": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string",
                    "example": "2023-05-14"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "company_model.CompanyInfo": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "联创"
                },
                "address": {
                    "type": "string",
                    "example": "启明学院亮胜楼八楼"
                },
                "business_license_number": {
                    "type": "string",
                    "example": "114514"
                },
                "manager_name": {
                    "type": "string",
                    "example": "汉堡"
                },
                "manager_tel": {
                    "type": "string",
                    "example": "1919810"
                },
                "name": {
                    "type": "string",
                    "example": "联小创在线科技有限公司"
                }
            }
        },
        "request.AdminCreateReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.AdminLoginReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "fnfunfunc"
                },
                "password": {
                    "type": "string",
                    "example": "1234"
                }
            }
        },
        "request.AllowCompaniesRegisterReq": {
            "type": "object",
            "properties": {
                "company_accounts": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "request.AllowCompaniesUpdateReq": {
            "type": "object",
            "properties": {
                "company_accounts": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "request.CompanyLoginReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.CompanyRegisterReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "联创"
                },
                "address": {
                    "type": "string",
                    "example": "启明学院亮胜楼八楼"
                },
                "business_license_number": {
                    "type": "string",
                    "example": "114514"
                },
                "manager_name": {
                    "type": "string",
                    "example": "汉堡"
                },
                "manager_tel": {
                    "type": "string",
                    "example": "1919810"
                },
                "name": {
                    "type": "string",
                    "example": "联小创在线科技有限公司"
                },
                "password": {
                    "type": "string",
                    "example": "uniquestudio"
                }
            }
        },
        "request.CompanyUpdateInfoReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "联创"
                },
                "address": {
                    "type": "string",
                    "example": "启明学院亮胜楼八楼"
                },
                "business_license_number": {
                    "type": "string",
                    "example": "114514"
                },
                "manager_name": {
                    "type": "string",
                    "example": "汉堡"
                },
                "manager_tel": {
                    "type": "string",
                    "example": "1919810"
                },
                "name": {
                    "type": "string",
                    "example": "联小创在线科技有限公司"
                }
            }
        },
        "request.CompanyUpdatePwdReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "1919810"
                }
            }
        },
        "request.CompanyUploadAdvtReq": {
            "type": "object",
            "properties": {
                "advertisementInfo": {
                    "$ref": "#/definitions/advertisement_model.AdvertisementInfo"
                }
            }
        },
        "response.AdminCreateResp": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                }
            }
        },
        "response.AdminLoginResp": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "fnfunfunc"
                },
                "token": {
                    "type": "string",
                    "example": "114514"
                }
            }
        },
        "response.CompanyGetInfoResp": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "联创"
                },
                "address": {
                    "type": "string",
                    "example": "启明学院亮胜楼八楼"
                },
                "business_license_number": {
                    "type": "string",
                    "example": "114514"
                },
                "manager_name": {
                    "type": "string",
                    "example": "汉堡"
                },
                "manager_tel": {
                    "type": "string",
                    "example": "1919810"
                },
                "name": {
                    "type": "string",
                    "example": "联小创在线科技有限公司"
                }
            }
        },
        "response.CompanyLoginResp": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "联创"
                },
                "token": {
                    "type": "string",
                    "example": "114514"
                }
            }
        },
        "response.GetAdvertisementsResp": {
            "type": "object",
            "properties": {
                "advertisement_infos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/advertisement_model.AdvertisementToBePreviewedInfo"
                    }
                }
            }
        },
        "response.GetCompaniesCountResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "response.GetCompaniesResp": {
            "type": "object",
            "properties": {
                "company_infos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/company_model.CompanyInfo"
                    }
                }
            }
        },
        "response.GetCompaniesToBeReviewedResp": {
            "type": "object",
            "properties": {
                "company_infos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/company_model.CompanyInfo"
                    }
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
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
