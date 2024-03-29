package echopb 

const (
EchoSwagger = `{
  "swagger": "2.0",
  "info": {
    "title": "echo.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/echo/api/v1/chat": {
      "post": {
        "operationId": "Chat",
        "responses": {
          "200": {
            "description": "A successful response.(streaming responses)",
            "schema": {
              "$ref": "#/x-stream-definitions/echopbEchoReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "description": " (streaming inputs)",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoRequest"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/echo/api/v1/echo": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/echopbEchoReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoRequest"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/echo/api/v1/sink": {
      "post": {
        "operationId": "Sink",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/echopbEchoReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "description": " (streaming inputs)",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoRequest"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/echo/api/v1/trico": {
      "post": {
        "operationId": "Trico",
        "responses": {
          "200": {
            "description": "A successful response.(streaming responses)",
            "schema": {
              "$ref": "#/x-stream-definitions/echopbEchoReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoRequest"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    }
  },
  "definitions": {
    "echopbEchoReply": {
      "type": "object",
      "properties": {
        "index": {
          "type": "integer",
          "format": "int32"
        },
        "msg": {
          "type": "string"
        }
      }
    },
    "echopbEchoRequest": {
      "type": "object",
      "properties": {
        "index": {
          "type": "integer",
          "format": "int32"
        },
        "msg": {
          "type": "string"
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "runtimeStreamError": {
      "type": "object",
      "properties": {
        "grpc_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "http_status": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    }
  },
  "x-stream-definitions": {
    "echopbEchoReply": {
      "type": "object",
      "properties": {
        "result": {
          "$ref": "#/definitions/echopbEchoReply"
        },
        "error": {
          "$ref": "#/definitions/runtimeStreamError"
        }
      },
      "title": "Stream result of echopbEchoReply"
    }
  }
}
`
)
