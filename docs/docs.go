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
        "/activities": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Takes an activity JSON and store in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activities"
                ],
                "summary": "Store a new activity",
                "parameters": [
                    {
                        "description": "Activity JSON",
                        "name": "activity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Activity"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Activity"
                        }
                    }
                }
            }
        },
        "/activities/{baby_id}/latest": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns all baby's activities.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activities"
                ],
                "summary": "Get all latest activities of baby",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search activity by baby_id",
                        "name": "baby_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Activity"
                        }
                    }
                }
            }
        },
        "/babies": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns all babies.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Babies"
                ],
                "summary": "Get all babies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Baby"
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
                "description": "Takes a baby JSON and store in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Babies"
                ],
                "summary": "Store a new baby",
                "parameters": [
                    {
                        "description": "Baby JSON",
                        "name": "baby",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Baby"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Baby"
                        }
                    }
                }
            }
        },
        "/babies/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns the baby whose id value matches the id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Babies"
                ],
                "summary": "Get single baby by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search baby by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Baby"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Activity": {
            "type": "object",
            "required": [
                "baby_id",
                "time",
                "type"
            ],
            "properties": {
                "baby_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "handler.Baby": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "dob": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "NFC Baby Tracker API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
