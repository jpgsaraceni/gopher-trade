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
            "name": "João Saraceni",
            "url": "https://www.linkedin.com/in/joaosaraceni/",
            "email": "jpgome@id.uff.br"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/jpgsaraceni/gopher-trade/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/exchanges": {
            "post": {
                "description": "Creates an exchange rate from and to specified currencies.\nNote that from-to currency pairs must be unique.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exchange"
                ],
                "summary": "Create a new exchange rate",
                "parameters": [
                    {
                        "description": "Exchange Info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/exchanges.CreateExchangeRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/exchanges.CreateExchangeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorPayload"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorPayload"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "exchanges.CreateExchangeRequest": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string",
                    "example": "USD"
                },
                "rate": {
                    "type": "string",
                    "example": "2.132"
                },
                "to": {
                    "type": "string",
                    "example": "COOLCOIN"
                }
            }
        },
        "exchanges.CreateExchangeResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "from": {
                    "type": "string",
                    "example": "USD"
                },
                "id": {
                    "type": "string",
                    "example": "2171f348-54b4-4a1e-8643-0972a3daf400"
                },
                "rate": {
                    "type": "string",
                    "example": "2.132"
                },
                "to": {
                    "type": "string",
                    "example": "COOLCOIN"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "responses.ErrorPayload": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Message for some error"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Gopher Trade API",
	Description:      "Gopher Trade is an api to get monetary exchange values.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}