package config

import "fmt"

const NODE_URL = "node.xelis.io"
const TESTNET_NODE_URL = "testnet-node.xelis.io"

var NODE_RPC = fmt.Sprintf("https://%s/json_rpc", NODE_URL)
var TESTNET_NODE_RPC = fmt.Sprintf("https://%s/json_rpc", TESTNET_NODE_URL)

var NODE_WS = fmt.Sprintf("wss://%s/json_rpc", NODE_URL)
var TESTNET_NODE_WS = fmt.Sprintf("wss://%s/json_rpc", TESTNET_NODE_URL)
