// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The BSC \u003c-\u003e EVM Compatible Swap API: provide swap service between BSC and EVM Compatible, which is based on https://github.com/binance-chain/eth-swap-ap",
    "title": "BSC \u003c-\u003e EVM Compatible Bridge API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.1.0"
  },
  "basePath": "/",
  "paths": {
    "/v1/erc-1155-swap-pairs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_1155_swap_pairs"
        ],
        "summary": "Gets a list of available ERC1155 swap pairs.",
        "operationId": "getErc1155SwapPairs",
        "parameters": [
          {
            "maximum": 10000,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 10000,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "enum": [
              "registration_ongoing",
              "registration_confirmed",
              "creation_tx_dry_run_failed",
              "creation_tx_created",
              "creation_tx_sent",
              "creation_tx_confirmed",
              "creation_tx_failed",
              "creation_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "description": "source chain id",
            "name": "src_chain_id",
            "in": "query"
          },
          {
            "type": "string",
            "description": "destination chain id",
            "name": "dst_chain_id",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "source token address",
            "name": "src_token_addr",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "destination token address",
            "name": "dst_token_addr",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc1155SwapPairs"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/erc-1155-swaps": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_1155_swaps"
        ],
        "summary": "Gets list of ERC1155 swap.",
        "operationId": "getErc1155Swaps",
        "parameters": [
          {
            "maximum": 1000,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 1000,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "address",
            "name": "sender",
            "in": "query",
            "required": true
          },
          {
            "enum": [
              "request_ongoing",
              "request_rejected",
              "request_confirmed",
              "fill_tx_dry_run_failed",
              "fill_tx_created",
              "fill_tx_sent",
              "fill_tx_confirmed",
              "fill_tx_failed",
              "fill_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{64}$",
            "type": "string",
            "description": "request_tx_hash",
            "name": "request_tx_hash",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc1155Swaps"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/erc-721-swap-pairs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_721_swap_pairs"
        ],
        "summary": "Gets a list of available ERC721 swap pairs.",
        "operationId": "getErc721SwapPairs",
        "parameters": [
          {
            "maximum": 10000,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 10000,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "enum": [
              "registration_ongoing",
              "registration_confirmed",
              "creation_tx_dry_run_failed",
              "creation_tx_created",
              "creation_tx_sent",
              "creation_tx_confirmed",
              "creation_tx_failed",
              "creation_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "description": "source chain id",
            "name": "src_chain_id",
            "in": "query"
          },
          {
            "type": "string",
            "description": "destination chain id",
            "name": "dst_chain_id",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "source token address",
            "name": "src_token_addr",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "destination token address",
            "name": "dst_token_addr",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc721SwapPairs"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/erc-721-swaps": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_721_swaps"
        ],
        "summary": "Gets list of ERC721 swap.",
        "operationId": "getErc721Swaps",
        "parameters": [
          {
            "maximum": 1000,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 1000,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "address",
            "name": "sender",
            "in": "query",
            "required": true
          },
          {
            "enum": [
              "request_ongoing",
              "request_rejected",
              "request_confirmed",
              "fill_tx_dry_run_failed",
              "fill_tx_created",
              "fill_tx_sent",
              "fill_tx_confirmed",
              "fill_tx_failed",
              "fill_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{64}$",
            "type": "string",
            "description": "request_tx_hash",
            "name": "request_tx_hash",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc721Swaps"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/info": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "svcInfo"
        ],
        "summary": "Gets service info",
        "operationId": "getInfo",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ServiceInfo"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Erc1155Swap": {
      "type": "object",
      "properties": {
        "amounts": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "fill_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "recipient": {
          "type": "string",
          "x-omitempty": false
        },
        "request_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sender": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "swap_direction": {
          "type": "string",
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc1155SwapPair": {
      "type": "object",
      "properties": {
        "available": {
          "type": "boolean",
          "x-omitempty": false
        },
        "create_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "register_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sponsor": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        },
        "uri": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc1155SwapPairs": {
      "type": "object",
      "properties": {
        "pairs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc1155SwapPair"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Erc1155Swaps": {
      "type": "object",
      "properties": {
        "erc_1155_swaps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc1155Swap"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Erc721Swap": {
      "type": "object",
      "properties": {
        "base_uri": {
          "type": "string",
          "x-omitempty": false
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "fill_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "recipient": {
          "type": "string",
          "x-omitempty": false
        },
        "request_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sender": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "swap_direction": {
          "type": "string",
          "x-omitempty": false
        },
        "token_id": {
          "type": "string",
          "x-omitempty": false
        },
        "token_uri": {
          "type": "string",
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc721SwapPair": {
      "type": "object",
      "properties": {
        "available": {
          "type": "boolean",
          "x-omitempty": false
        },
        "base_uri": {
          "type": "string",
          "x-omitempty": false
        },
        "create_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "register_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sponsor": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "symbol": {
          "type": "string",
          "x-omitempty": false,
          "example": "USDT"
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc721SwapPairs": {
      "type": "object",
      "properties": {
        "pairs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc721SwapPair"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Erc721Swaps": {
      "type": "object",
      "properties": {
        "erc_721_swaps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc721Swap"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ServiceInfo": {
      "type": "object",
      "properties": {
        "bsc_chain_id": {
          "type": "integer",
          "x-omitempty": false
        },
        "bsc_erc_1155_swap_agent": {
          "type": "string",
          "x-omitempty": false
        },
        "bsc_erc_721_swap_agent": {
          "type": "string",
          "x-omitempty": false
        },
        "eth_chain_id": {
          "type": "integer",
          "x-omitempty": false
        },
        "eth_erc_1155_swap_agent": {
          "type": "string",
          "x-omitempty": false
        },
        "eth_erc_721_swap_agent": {
          "type": "string",
          "x-omitempty": false
        }
      }
    }
  },
  "tags": [
    {
      "description": "Erc721Swap Pair list",
      "name": "erc_721_swap_pairs"
    },
    {
      "description": "Erc721Swap list",
      "name": "erc_721_swaps"
    },
    {
      "description": "Erc1155Swap Pair list",
      "name": "erc_1155_swap_pairs"
    },
    {
      "description": "Erc1155Swap list",
      "name": "erc_1155_swaps"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The BSC \u003c-\u003e EVM Compatible Swap API: provide swap service between BSC and EVM Compatible, which is based on https://github.com/binance-chain/eth-swap-ap",
    "title": "BSC \u003c-\u003e EVM Compatible Bridge API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.1.0"
  },
  "basePath": "/",
  "paths": {
    "/v1/erc-1155-swap-pairs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_1155_swap_pairs"
        ],
        "summary": "Gets a list of available ERC1155 swap pairs.",
        "operationId": "getErc1155SwapPairs",
        "parameters": [
          {
            "maximum": 10000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 10000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "enum": [
              "registration_ongoing",
              "registration_confirmed",
              "creation_tx_dry_run_failed",
              "creation_tx_created",
              "creation_tx_sent",
              "creation_tx_confirmed",
              "creation_tx_failed",
              "creation_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "description": "source chain id",
            "name": "src_chain_id",
            "in": "query"
          },
          {
            "type": "string",
            "description": "destination chain id",
            "name": "dst_chain_id",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "source token address",
            "name": "src_token_addr",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "destination token address",
            "name": "dst_token_addr",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc1155SwapPairs"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/erc-1155-swaps": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_1155_swaps"
        ],
        "summary": "Gets list of ERC1155 swap.",
        "operationId": "getErc1155Swaps",
        "parameters": [
          {
            "maximum": 1000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 1000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "address",
            "name": "sender",
            "in": "query",
            "required": true
          },
          {
            "enum": [
              "request_ongoing",
              "request_rejected",
              "request_confirmed",
              "fill_tx_dry_run_failed",
              "fill_tx_created",
              "fill_tx_sent",
              "fill_tx_confirmed",
              "fill_tx_failed",
              "fill_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{64}$",
            "type": "string",
            "description": "request_tx_hash",
            "name": "request_tx_hash",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc1155Swaps"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/erc-721-swap-pairs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_721_swap_pairs"
        ],
        "summary": "Gets a list of available ERC721 swap pairs.",
        "operationId": "getErc721SwapPairs",
        "parameters": [
          {
            "maximum": 10000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 10000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "enum": [
              "registration_ongoing",
              "registration_confirmed",
              "creation_tx_dry_run_failed",
              "creation_tx_created",
              "creation_tx_sent",
              "creation_tx_confirmed",
              "creation_tx_failed",
              "creation_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "description": "source chain id",
            "name": "src_chain_id",
            "in": "query"
          },
          {
            "type": "string",
            "description": "destination chain id",
            "name": "dst_chain_id",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "source token address",
            "name": "src_token_addr",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "destination token address",
            "name": "dst_token_addr",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc721SwapPairs"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/erc-721-swaps": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "erc_721_swaps"
        ],
        "summary": "Gets list of ERC721 swap.",
        "operationId": "getErc721Swaps",
        "parameters": [
          {
            "maximum": 1000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 0,
            "description": "offset",
            "name": "offset",
            "in": "query"
          },
          {
            "maximum": 1000,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "default": 100,
            "description": "limit",
            "name": "limit",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{40}$",
            "type": "string",
            "description": "address",
            "name": "sender",
            "in": "query",
            "required": true
          },
          {
            "enum": [
              "request_ongoing",
              "request_rejected",
              "request_confirmed",
              "fill_tx_dry_run_failed",
              "fill_tx_created",
              "fill_tx_sent",
              "fill_tx_confirmed",
              "fill_tx_failed",
              "fill_tx_missing"
            ],
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          },
          {
            "pattern": "^(0x)[0-9A-Fa-f]{64}$",
            "type": "string",
            "description": "request_tx_hash",
            "name": "request_tx_hash",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Erc721Swaps"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/info": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "svcInfo"
        ],
        "summary": "Gets service info",
        "operationId": "getInfo",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ServiceInfo"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Erc1155Swap": {
      "type": "object",
      "properties": {
        "amounts": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "fill_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "recipient": {
          "type": "string",
          "x-omitempty": false
        },
        "request_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sender": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "swap_direction": {
          "type": "string",
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc1155SwapPair": {
      "type": "object",
      "properties": {
        "available": {
          "type": "boolean",
          "x-omitempty": false
        },
        "create_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "register_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sponsor": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        },
        "uri": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc1155SwapPairs": {
      "type": "object",
      "properties": {
        "pairs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc1155SwapPair"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Erc1155Swaps": {
      "type": "object",
      "properties": {
        "erc_1155_swaps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc1155Swap"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Erc721Swap": {
      "type": "object",
      "properties": {
        "base_uri": {
          "type": "string",
          "x-omitempty": false
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "fill_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "recipient": {
          "type": "string",
          "x-omitempty": false
        },
        "request_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sender": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "swap_direction": {
          "type": "string",
          "x-omitempty": false
        },
        "token_id": {
          "type": "string",
          "x-omitempty": false
        },
        "token_uri": {
          "type": "string",
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc721SwapPair": {
      "type": "object",
      "properties": {
        "available": {
          "type": "boolean",
          "x-omitempty": false
        },
        "base_uri": {
          "type": "string",
          "x-omitempty": false
        },
        "create_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "created_at": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "dst_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "register_tx_hash": {
          "type": "string",
          "x-omitempty": false
        },
        "sponsor": {
          "type": "string",
          "x-omitempty": false
        },
        "src_chain_id": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_addr": {
          "type": "string",
          "x-omitempty": false
        },
        "src_token_name": {
          "type": "string",
          "x-omitempty": false
        },
        "state": {
          "type": "string",
          "x-omitempty": false
        },
        "symbol": {
          "type": "string",
          "x-omitempty": false,
          "example": "USDT"
        },
        "updated_at": {
          "type": "string",
          "x-omitempty": false
        }
      }
    },
    "Erc721SwapPairs": {
      "type": "object",
      "properties": {
        "pairs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc721SwapPair"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Erc721Swaps": {
      "type": "object",
      "properties": {
        "erc_721_swaps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Erc721Swap"
          },
          "x-omitempty": false
        },
        "total": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ServiceInfo": {
      "type": "object",
      "properties": {
        "bsc_chain_id": {
          "type": "integer",
          "x-omitempty": false
        },
        "bsc_erc_1155_swap_agent": {
          "type": "string",
          "x-omitempty": false
        },
        "bsc_erc_721_swap_agent": {
          "type": "string",
          "x-omitempty": false
        },
        "eth_chain_id": {
          "type": "integer",
          "x-omitempty": false
        },
        "eth_erc_1155_swap_agent": {
          "type": "string",
          "x-omitempty": false
        },
        "eth_erc_721_swap_agent": {
          "type": "string",
          "x-omitempty": false
        }
      }
    }
  },
  "tags": [
    {
      "description": "Erc721Swap Pair list",
      "name": "erc_721_swap_pairs"
    },
    {
      "description": "Erc721Swap list",
      "name": "erc_721_swaps"
    },
    {
      "description": "Erc1155Swap Pair list",
      "name": "erc_1155_swap_pairs"
    },
    {
      "description": "Erc1155Swap list",
      "name": "erc_1155_swaps"
    }
  ]
}`))
}
