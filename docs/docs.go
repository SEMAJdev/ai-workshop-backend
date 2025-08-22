package docs

// This file provides a hand-written Swagger spec without using swag CLI.

import (
	"encoding/json"

	"github.com/swaggo/swag"
)

// SwaggerInfo holds basic API metadata
var SwaggerInfo = struct {
	Title       string
	Description string
	Version     string
	BasePath    string
	Schemes     []string
}{
	Title:       "Workshop Backend API",
	Description: "Authentication and Profile APIs",
	Version:     "1.0.0",
	BasePath:    "/",
	Schemes:     []string{"http"},
}

// ReadDoc implements fiber-swagger expected function to fetch the spec.
func ReadDoc() string {
	// Build minimal OpenAPI 3.0 JSON
	spec := map[string]any{
		"openapi": "3.0.0",
		"info": map[string]any{
			"title":       SwaggerInfo.Title,
			"description": SwaggerInfo.Description,
			"version":     SwaggerInfo.Version,
		},
		"servers": []map[string]any{{
			"url": SwaggerInfo.BasePath,
		}},
		"paths": map[string]any{
			"/api/login": map[string]any{
				"post": map[string]any{
					"summary": "Login with email and password",
					"requestBody": map[string]any{
						"required": true,
						"content": map[string]any{
							"application/json": map[string]any{
								"schema": map[string]any{
									"type": "object",
									"properties": map[string]any{
										"email":    map[string]any{"type": "string", "example": "somchai@example.com"},
										"password": map[string]any{"type": "string", "example": "password123"},
									},
									"required": []string{"email", "password"},
								},
							},
						},
					},
					"responses": map[string]any{
						"200": map[string]any{
							"description": "OK",
							"content": map[string]any{
								"application/json": map[string]any{
									"schema": map[string]any{
										"type": "object",
										"properties": map[string]any{
											"token": map[string]any{"type": "string"},
											"user":  map[string]any{"$ref": "#/components/schemas/PublicProfile"},
										},
									},
								},
							},
						},
						"401": map[string]any{"description": "Unauthorized"},
					},
				},
			},
			"/api/me": map[string]any{
				"get": map[string]any{
					"summary":  "Get current user profile",
					"security": []map[string]any{{"bearerAuth": []any{}}},
					"responses": map[string]any{
						"200": map[string]any{
							"description": "OK",
							"content": map[string]any{
								"application/json": map[string]any{
									"schema": map[string]any{"$ref": "#/components/schemas/PublicProfile"},
								},
							},
						},
						"401": map[string]any{"description": "Unauthorized"},
					},
				},
				"put": map[string]any{
					"summary":  "Update current user profile",
					"security": []map[string]any{{"bearerAuth": []any{}}},
					"requestBody": map[string]any{
						"required": true,
						"content": map[string]any{
							"application/json": map[string]any{
								"schema": map[string]any{"$ref": "#/components/schemas/UpdateProfileInput"},
							},
						},
					},
					"responses": map[string]any{
						"200": map[string]any{
							"description": "OK",
							"content": map[string]any{
								"application/json": map[string]any{
									"schema": map[string]any{"$ref": "#/components/schemas/PublicProfile"},
								},
							},
						},
						"401": map[string]any{"description": "Unauthorized"},
					},
				},
			},
		},
		"components": map[string]any{
			"securitySchemes": map[string]any{
				"bearerAuth": map[string]any{
					"type":         "http",
					"scheme":       "bearer",
					"bearerFormat": "JWT",
				},
			},
			"schemas": map[string]any{
				"PublicProfile": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"id":              map[string]any{"type": "integer", "format": "int64"},
						"email":           map[string]any{"type": "string"},
						"firstName":       map[string]any{"type": "string"},
						"lastName":        map[string]any{"type": "string"},
						"phone":           map[string]any{"type": "string"},
						"memberCode":      map[string]any{"type": "string"},
						"membershipLevel": map[string]any{"type": "string"},
						"points":          map[string]any{"type": "integer"},
						"joinedAt":        map[string]any{"type": "string", "format": "date-time"},
					},
				},
				"UpdateProfileInput": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"firstName": map[string]any{"type": "string"},
						"lastName":  map[string]any{"type": "string"},
						"phone":     map[string]any{"type": "string"},
					},
					"required": []string{"firstName", "lastName"},
				},
			},
		},
	}
	b, _ := json.Marshal(spec)
	return string(b)
}

// Register the ReadDoc with swag so fiber-swagger can find it
type swaggerDoc struct{}

func (swaggerDoc) ReadDoc() string { return ReadDoc() }

func init() { swag.Register(swag.Name, &swaggerDoc{}) }
