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
            "name": "Muhammad Talha",
            "email": "talhach891@gmail.com"
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
        "/partners": {
            "get": {
                "description": "This API can be used to retrieve best possible matched partners w.r.t distance and rating,",
                "summary": "Get best possible matched partners",
                "parameters": [
                    {
                        "enum": [
                            "wood",
                            "carpet",
                            "tiles"
                        ],
                        "type": "string",
                        "description": "Floor material",
                        "name": "material",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Latitude(example: 53.544422)",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Longitude(example: 10.0011)",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_server_apihandlers.partnersListResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/forms.BasicResponse"
                        }
                    }
                }
            }
        },
        "/partners/{id}": {
            "get": {
                "description": "This API returns partner details.",
                "summary": "Get partner's details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Partner Id(example 272)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.com_taalhach_aroundhome-challennge_internal_server_apihandlers.partnerDetailsResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/forms.BasicResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dbutils.PartnerListItem": {
            "type": "object",
            "properties": {
                "distance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "materials": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "radius_in_meters": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "forms.BasicResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github.com_taalhach_aroundhome-challennge_internal_server_apihandlers.partnerDetailsResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "partner": {
                    "$ref": "#/definitions/dbutils.PartnerListItem"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "github.com_taalhach_aroundhome-challennge_internal_server_apihandlers.partnersListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dbutils.PartnerListItem"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pages": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "internal_server_apihandlers.partnerDetailsResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "partner": {
                    "$ref": "#/definitions/dbutils.PartnerListItem"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "internal_server_apihandlers.partnersListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dbutils.PartnerListItem"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pages": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "aroundhome-challennge API docs",
	Description:      "aroundhome's code aroundhome API specs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
