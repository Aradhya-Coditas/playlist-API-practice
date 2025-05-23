// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/Playlist/create": {
            "post": {
                "description": "Creates a new playlist for a user with specified songs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Playlists"
                ],
                "summary": "Create a new playlist",
                "parameters": [
                    {
                        "description": "Playlist Creation Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BFFPlaylistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created playlist",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request - Invalid input or validation errors",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/api/watchlist/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete Watchlist API for deleting the watchlist created by the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete Watchlist"
                ],
                "summary": "Delete Watchlist API",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123456789",
                        "description": "Unique request identifier",
                        "name": "xRequestId",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "PKQ1.180904.001",
                        "description": "Unique device identifier",
                        "name": "deviceId",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1.0.0",
                        "description": "Current app version",
                        "name": "appVersion",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "MOB",
                        "description": "Source (MOB or WEB)",
                        "name": "source",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bypass (AUTOMATION or CHART)",
                        "name": "bypass",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "default": "ba6eb330-4f7f-11eb-a2fb-67c34e9ac07c",
                        "description": "Unique appInstall identifier",
                        "name": "appInstallId",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "BrokerAppName/3.3.6 (OnePlus ONEPLUS A6010; Android 11 SDK30",
                        "description": "userAgent",
                        "name": "userAgent",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1700839140000",
                        "description": "device current day epoch milliseconds timestamp",
                        "name": "timestamp",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "DeleteWatchlistRequest JSON",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BFFDeleteWatchlistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BFFDeleteWatchlistResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found: User not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorAPIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorAPIResponse"
                        }
                    }
                }
            }
        },
        "/api/watchlist/get": {
            "post": {
                "description": "Returns user's custom and broker's predefined watchlists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Watchlist"
                ],
                "summary": "Get user and broker watchlists",
                "parameters": [
                    {
                        "description": "Request payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BFFGetWatchlistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BFFWatchlistResponse"
                        }
                    },
                    "204": {
                        "description": "No content found"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorAPIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorAPIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorAPIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BFFDeleteWatchlistRequest": {
            "type": "object",
            "required": [
                "watchlistId"
            ],
            "properties": {
                "watchlistId": {
                    "type": "integer",
                    "example": 30
                }
            }
        },
        "models.BFFDeleteWatchlistResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "models.BFFGetWatchlistRequest": {
            "type": "object",
            "required": [
                "brokerId",
                "userId"
            ],
            "properties": {
                "brokerId": {
                    "type": "integer",
                    "example": 1
                },
                "userId": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "models.BFFPlaylistRequest": {
            "type": "object",
            "required": [
                "name",
                "user_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "song_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.BFFPredefine": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "watchlistName": {
                    "type": "string"
                }
            }
        },
        "models.BFFUserdefine": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "watchlistName": {
                    "type": "string"
                }
            }
        },
        "models.BFFWatchlistResponse": {
            "type": "object",
            "properties": {
                "predefine": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BFFPredefine"
                    }
                },
                "userdefine": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BFFUserdefine"
                    }
                }
            }
        },
        "models.ErrorAPIResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ErrorMessage"
                    }
                }
            }
        },
        "models.ErrorMessage": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header",
            "x-extension-openapi": "{\"example\": \"value on a json format\"}"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "omnenest-backend",
	Description:      "Omnenest backend for watchlist micro-service (Middleware layer).",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
