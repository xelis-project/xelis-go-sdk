package xelis

import "fmt"

const NODE_URL = "node.xelis.io"
const TESTNET_NODE_URL = "testnet-node.xelis.io"
const DEV_NODE_URL = "dev-node.xelis.io"

var NODE_RPC = fmt.Sprintf("https://%s/json_rpc", NODE_URL)
var TESTNET_NODE_RPC = fmt.Sprintf("https://%s/json_rpc", TESTNET_NODE_URL)
var DEV_NODE_RPC = fmt.Sprintf("https://%s/json_rpc", DEV_NODE_URL)

var NODE_WS = fmt.Sprintf("wss://%s/ws", NODE_URL)
var TESTNET_NODE_WS = fmt.Sprintf("wss://%s/ws", TESTNET_NODE_URL)
var DEV_NODE_WS = fmt.Sprintf("wss://%s/ws", DEV_NODE_URL)
