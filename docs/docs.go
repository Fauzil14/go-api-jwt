// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/age-rating-categories": {
            "get": {
                "description": "Get List of Age Rating Category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get All Age Rating Category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AgeRatingCategory"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get List of Age Rating Category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Create Age Rating Category",
                "parameters": [
                    {
                        "description": "the body to create new age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AgeRatingCategoryInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorizattion",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            }
        },
        "/age-rating-categories/{id}": {
            "get": {
                "description": "Get one Age Rating Category by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get an Age Rating Category by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Rating Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one Age Rating Category by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Delete an Age Rating Category by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Rating Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update one Age Rating Category by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Update an Age Rating Category by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Rating Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update Age Rating Category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AgeRatingCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            }
        },
        "/age-rating-categories/{id}/movies": {
            "get": {
                "description": "Get all movies of Age Rating Category by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get movies by Age Rating Category by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Age Rating Category id",
                        "name": "id",
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
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logging in to get jwt token for authorization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login as an User",
                "parameters": [
                    {
                        "description": "the body to login",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "Get List of Movie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get All Movie",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Get List of Movie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Create Movie",
                "parameters": [
                    {
                        "description": "the body to create new Movie",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.MovieInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "description": "Get one Movie by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get a Movie by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete one Movie by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Delete a Movie by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update one Movie by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Update a Movie by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update Movie",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.MovieInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register user from public access",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new User",
                "parameters": [
                    {
                        "description": "the body to Register",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AgeRatingCategoryInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.MovieInput": {
            "type": "object",
            "properties": {
                "age_rating_category_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "controllers.RegisterInput": {
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
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.AgeRatingCategory": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "age_rating_category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
