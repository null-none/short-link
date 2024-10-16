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
        "/short-url": {
            "post": {
                "description": "Create new hash:url couple from given url",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "ShortUrl"
                ],
                "summary": "Creates the md5 hash of given URL string and stores it in DB",
                "parameters": [
                    {
                        "description": "Short Url Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ShortUrlReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ShortUrl"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/short-url/{id}": {
            "get": {
                "description": "get URL string from given hash of it",
                "tags": [
                    "ShortUrl"
                ],
                "summary": "Get URL string of a hash",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hash String",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ShortUrl"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ShortUrl": {
            "type": "object",
            "required": [
                "hash",
                "url"
            ],
            "properties": {
                "hash": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.ShortUrlReq": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:80",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Shortener API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
