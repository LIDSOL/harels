package main

import (
	"encoding/json"
)

type Message struct {
	Jsonrpc string `json:"jsonrpc"`
}

type RequestMessage struct {
	Message
	ID     interface{}     `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}

type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseMessage struct {
	Message
	ID     interface{}    `json:"id"`
	Result interface{}    `json:"result,omitempty"`
	Error  *ResponseError `json:"error,omitempty"`
}

type NotificationMessage struct {
	Message
	Method string          `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}

const (
	// Defined by JSON-RPC
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603

	// JSON-RPC reserved error range
	JSONRPCReservedErrorRangeStart = -32099
	ServerErrorStart               = JSONRPCReservedErrorRangeStart // Deprecated

	ServerNotInitialized = -32002
	UnknownErrorCode     = -32001

	JSONRPCReservedErrorRangeEnd = -32000
	ServerErrorEnd               = JSONRPCReservedErrorRangeEnd // Deprecated

	// LSP reserved error range
	LSPReservedErrorRangeStart = -32899

	RequestFailed    = -32803
	ServerCancelled  = -32802
	ContentModified  = -32801
	RequestCancelled = -32800

	LSPReservedErrorRangeEnd = -32800
)
