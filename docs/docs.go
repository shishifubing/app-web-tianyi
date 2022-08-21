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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/database/migrate": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "migrate databases",
                "tags": [
                    "database"
                ],
                "summary": "migrate databases",
                "operationId": "database-migrate",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.Success"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/projects": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get projects",
                "tags": [
                    "project"
                ],
                "summary": "get projects",
                "operationId": "get-projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search by path",
                        "name": "path",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Project"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create a project",
                "tags": [
                    "project"
                ],
                "summary": "create a project",
                "operationId": "create-project",
                "parameters": [
                    {
                        "description": "Project body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ProjectRequestCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Project"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/projects/{project_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get a project",
                "tags": [
                    "project"
                ],
                "summary": "get a project",
                "operationId": "get-project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Project"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update a project",
                "tags": [
                    "project"
                ],
                "summary": "update a project",
                "operationId": "update-project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Project body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Project"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/projects/{project_id}/branches": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all branches",
                "tags": [
                    "branch"
                ],
                "summary": "get all project branches",
                "operationId": "get-branches",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Branch"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create a branch",
                "tags": [
                    "branch"
                ],
                "summary": "create a branch",
                "operationId": "create-branch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "creation request body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RequestBranchCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Branch"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/projects/{project_id}/branches/{branch_name}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get a project branch",
                "tags": [
                    "branch"
                ],
                "summary": "get a project branch",
                "operationId": "get-branch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "branch name",
                        "name": "branch_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Branch"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update a branch",
                "tags": [
                    "branch"
                ],
                "summary": "update a branch",
                "operationId": "update-branch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "branch name",
                        "name": "branch_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Branch"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/projects/{project_id}/branches/{branch_name}/pipelines/{pipeline_name}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create a pipeline",
                "tags": [
                    "pipeline"
                ],
                "summary": "create a pipeline",
                "operationId": "create-pipeline",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "branch name",
                        "name": "branch_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "pipeline name",
                        "name": "pipeline_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all users",
                "tags": [
                    "user"
                ],
                "summary": "get all users",
                "operationId": "get-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create a user",
                "tags": [
                    "user"
                ],
                "summary": "create a user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "User body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UserRequestCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/users/login": {
            "get": {
                "description": "login",
                "tags": [
                    "user"
                ],
                "summary": "login",
                "operationId": "login-user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.UserLogin"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/users/user/{user_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get a user",
                "tags": [
                    "user"
                ],
                "summary": "get a user",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ProjectRequestCreate": {
            "type": "object",
            "required": [
                "name",
                "source"
            ],
            "properties": {
                "default_branch": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "controller.RequestBranchCreate": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 3
                }
            }
        },
        "controller.UserRequestCreate": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 3
                }
            }
        },
        "entity.Branch": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "config": {
                    "$ref": "#/definitions/entity.PipelineConfig"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 1
                },
                "project_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Capability": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 3
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Permission": {
            "type": "object",
            "required": [
                "name",
                "path"
            ],
            "properties": {
                "capabilities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Capability"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 3
                },
                "path": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.PipelineConfig": {
            "type": "object",
            "properties": {
                "jobs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PipelineConfigJob"
                    }
                },
                "pipelines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PipelineConfigPipeline"
                    }
                }
            }
        },
        "entity.PipelineConfigJob": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "query": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "request_type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "entity.PipelineConfigPipeline": {
            "type": "object",
            "properties": {
                "jobs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.PipelineConfigPipelineJob"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.PipelineConfigPipelineJob": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.Project": {
            "type": "object",
            "required": [
                "default_branch",
                "name",
                "source"
            ],
            "properties": {
                "branches": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Branch"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "default_branch": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 1
                },
                "deleted_at": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "description": "display name",
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 1
                },
                "namespace_id": {
                    "description": "foreign key for the parent namespace",
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Role": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 256,
                    "minLength": 3
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Permission"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "required": [
                "email",
                "username"
            ],
            "properties": {
                "admin": {
                    "type": "boolean"
                },
                "bio": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "roles": {
                    "description": "\"many to many\" association\nhttps://gorm.io/docs/many_to_many.html",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Role"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "pkg.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string",
                    "example": "fail"
                },
                "type": {
                    "type": "string",
                    "example": "*fiber.Error"
                }
            }
        },
        "presenter.Success": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "presenter.UserLogin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "sql.NullTime": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "\"Type 'Bearer TOKEN' to correctly set the API Key\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
