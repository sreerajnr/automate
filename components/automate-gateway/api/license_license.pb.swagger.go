package api

func init() {
	Swagger.Add("license_license", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/license/license.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/license/apply": {
      "post": {
        "operationId": "ApplyLicense",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.license.ApplyLicenseResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.license.ApplyLicenseReq"
            }
          }
        ],
        "tags": [
          "License"
        ]
      }
    },
    "/api/v0/license/request": {
      "post": {
        "operationId": "RequestLicense",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.license.RequestLicenseResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.license.RequestLicenseReq"
            }
          }
        ],
        "tags": [
          "License"
        ]
      }
    },
    "/api/v0/license/status": {
      "get": {
        "operationId": "GetStatus",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.license.GetStatusResp"
            }
          }
        },
        "tags": [
          "License"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.license.ApplyLicenseReq": {
      "type": "object",
      "properties": {
        "license": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.license.ApplyLicenseResp": {
      "type": "object",
      "properties": {
        "status": {
          "$ref": "#/definitions/chef.automate.api.license.GetStatusResp"
        }
      }
    },
    "chef.automate.api.license.GetStatusResp": {
      "type": "object",
      "properties": {
        "license_id": {
          "type": "string"
        },
        "configured_at": {
          "type": "string",
          "format": "date-time"
        },
        "licensed_period": {
          "$ref": "#/definitions/chef.automate.api.license.GetStatusResp.DateRange"
        },
        "customer_name": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.license.GetStatusResp.DateRange": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string",
          "format": "date-time"
        },
        "end": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "chef.automate.api.license.RequestLicenseReq": {
      "type": "object",
      "properties": {
        "first_name": {
          "type": "string"
        },
        "last_name": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "gdpr_agree": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "chef.automate.api.license.RequestLicenseResp": {
      "type": "object",
      "properties": {
        "license": {
          "type": "string"
        },
        "status": {
          "$ref": "#/definitions/chef.automate.api.license.GetStatusResp"
        }
      }
    }
  }
}
`)
}
