// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/register": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RequestRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ResponseRegister"
                        }
                    },
                    "409": {
                        "description": "Duplicate record",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Form validation error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.RequestRegister": {
            "type": "object",
            "required": [
                "name",
                "password",
                "username"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 100
                }
            }
        },
        "domain.ResponseRegister": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.ErrorMessageResponse": {
            "type": "object",
            "properties": {
                "attribute": {},
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/response.ErrorMessageResponse"
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
