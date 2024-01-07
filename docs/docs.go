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
        "/v1/auth/login": {
            "post": {
                "description": "Login account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    }
                }
            }
        },
        "/v1/auth/logout": {
            "get": {
                "description": "logout user from the service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/v1/auth/signup": {
            "post": {
                "description": "Register account for admin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Sign Up",
                        "name": "sign_up",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpResponse"
                        }
                    }
                }
            }
        },
        "/v1/common/healthcheck": {
            "get": {
                "description": "health check api server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/oauth2/login": {
            "get": {
                "description": "Auth0 Login form.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "get information for users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.UserInfo"
                        }
                    }
                }
            },
            "put": {
                "description": "update information for users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "description": "Update User",
                        "name": "UserInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        },
        "/v1/users/image": {
            "put": {
                "description": "update information for users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OK"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponse": {
            "type": "object",
            "required": [
                "email",
                "success",
                "token"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "domain.OK": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "phone_number"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpResponse": {
            "type": "object",
            "required": [
                "msg",
                "success"
            ],
            "properties": {
                "msg": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "domain.UserInfo": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "id",
                "image",
                "last_name",
                "phone_number",
                "role",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.UserUpdate": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "id",
                "last_name",
                "phone_number",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swiftcart API Documentation",
	Description:      "This is a documentation for the Swiftcart API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
