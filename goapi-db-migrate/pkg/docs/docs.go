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
        "/todos": {
            "get": {
                "description": "You can filter all existing todos by listing them.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "List all existing todos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filter the text based value (ex: term=dosomething)",
                        "name": "term",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "filter the status based value (ex: completed=true)",
                        "name": "completed",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Go to a specific page number. Start with 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size for the data",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Page order. Eg: text desc,createdAt desc",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagdto.ResponseWithPage"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/swagger.TodoSampleListData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error500"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Add a new todo",
                "parameters": [
                    {
                        "description": "Todo Data",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.CreateTodoFrom"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagdto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/swagger.TodoSampleData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagdto.Error422"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/swagger.ErrCreateSampleData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error500"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Get a specific todo by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get a todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
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
                                    "$ref": "#/definitions/swagdto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/swagger.TodoSampleData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error500"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a specific todo by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Delete a todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error404"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error500"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a specific todo status by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Update a todo status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo Status Data",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.UpdateTodoStatusForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagdto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/swagger.TodoSampleData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error400"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error404"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagdto.Error422"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/swagger.ErrUpdateSampleData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagdto.Error500"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "swagdto.Error400": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/swagdto.ErrorData400"
                },
                "requestId": {
                    "type": "string",
                    "example": "3b6272b9-1ef1-45e0"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "swagdto.Error404": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/swagdto.ErrorData404"
                },
                "requestId": {
                    "type": "string",
                    "example": "3b6272b9-1ef1-45e0"
                },
                "status": {
                    "type": "integer",
                    "example": 404
                }
            }
        },
        "swagdto.Error422": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/swagdto.ErrorData422"
                },
                "requestId": {
                    "type": "string",
                    "example": "3b6272b9-1ef1-45e0"
                },
                "status": {
                    "type": "integer",
                    "example": 422
                }
            }
        },
        "swagdto.Error500": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/swagdto.ErrorData500"
                },
                "requestId": {
                    "type": "string",
                    "example": "3b6272b9-1ef1-45e0"
                },
                "status": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "swagdto.ErrorData400": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "400"
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "swagdto.ErrorData404": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "404"
                },
                "message": {
                    "type": "string",
                    "example": "Not Found"
                }
            }
        },
        "swagdto.ErrorData422": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "422"
                },
                "details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/swagdto.ErrorDetail"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "invalid data see details"
                }
            }
        },
        "swagdto.ErrorData500": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "500"
                },
                "message": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "swagdto.ErrorDetail": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "name field is required"
                },
                "target": {
                    "type": "string",
                    "example": "name"
                }
            }
        },
        "swagdto.PagingResult": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 20
                },
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "nextPage": {
                    "type": "integer",
                    "example": 2
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "prevPage": {
                    "type": "integer",
                    "example": 0
                },
                "totalPage": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "swagdto.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "requestId": {
                    "type": "string",
                    "example": "3b6272b9-1ef1-45e0"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "swagdto.ResponseWithPage": {
            "type": "object",
            "properties": {
                "_pagination": {
                    "$ref": "#/definitions/swagdto.PagingResult"
                },
                "data": {},
                "requestId": {
                    "type": "string",
                    "example": "3b6272b9-1ef1-45e0"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "swagger.CreateTodoFrom": {
            "type": "object",
            "properties": {
                "text": {
                    "description": "Required: true",
                    "type": "string",
                    "example": "do something"
                }
            }
        },
        "swagger.ErrCreateSampleData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "422"
                },
                "details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/swagger.ErrorDetailCreate"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "invalid data see details"
                }
            }
        },
        "swagger.ErrUpdateSampleData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "422"
                },
                "details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/swagger.ErrorDetailUpdate"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "invalid data see details"
                }
            }
        },
        "swagger.ErrorDetailCreate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "text field is required"
                },
                "target": {
                    "type": "string",
                    "example": "text"
                }
            }
        },
        "swagger.ErrorDetailUpdate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "completed field is required"
                },
                "target": {
                    "type": "string",
                    "example": "completed"
                }
            }
        },
        "swagger.TodoRepsonse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "id": {
                    "type": "string",
                    "example": "bfbc2a69-9825-4a0e-a8d6-ffb985dc719c"
                },
                "text": {
                    "type": "string",
                    "example": "do something"
                }
            }
        },
        "swagger.TodoSampleData": {
            "type": "object",
            "properties": {
                "todo": {
                    "$ref": "#/definitions/swagger.TodoRepsonse"
                }
            }
        },
        "swagger.TodoSampleListData": {
            "type": "object",
            "properties": {
                "todos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/swagger.TodoRepsonse"
                    }
                }
            }
        },
        "swagger.UpdateTodoStatusForm": {
            "type": "object",
            "properties": {
                "completed": {
                    "description": "Required: true",
                    "type": "boolean"
                }
            }
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
