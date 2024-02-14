package lib

import "encoding/json"

type RPCMethod string
type RPCEvent string

type RPCRequest struct {
	ID      int64       `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Method  RPCMethod   `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	ID     int64           `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
