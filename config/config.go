package config

import "fmt"

const DAEMON_RPC_PORT = 8080
const WALLET_RPC_PORT = 8081
const XSWD_PORT = 44325

const MAINNET_NODE_URL = "node.xelis.io"
const TESTNET_NODE_URL = "testnet-node.xelis.io"

const TESTNET_WALLET = "xet:6eadzwf5xdacts6fs4y3csmnsmy4mcxewqt3xyygwfx0hm0tm32sqxdy9zk"
const MAINNET_WALLET = "xel:vs3mfyywt0fjys0rgslue7mm4wr23xdgejsjk0ld7f2kxng4d4nqqnkdufz"

var LOCAL_NODE_URL = fmt.Sprintf("127.0.0.1:%d", DAEMON_RPC_PORT)

var MAINNET_NODE_RPC = fmt.Sprintf("https://%s/json_rpc", MAINNET_NODE_URL)
var TESTNET_NODE_RPC = fmt.Sprintf("https://%s/json_rpc", TESTNET_NODE_URL)
var LOCAL_NODE_RPC = fmt.Sprintf("http://%s/json_rpc", LOCAL_NODE_URL)

var MAINNET_NODE_WS = fmt.Sprintf("wss://%s/json_rpc", MAINNET_NODE_URL)
var TESTNET_NODE_WS = fmt.Sprintf("wss://%s/json_rpc", TESTNET_NODE_URL)
var LOCAL_NODE_WS = fmt.Sprintf("wss://%s/json_rpc", LOCAL_NODE_URL)

var TESTNET_NODE_GETWORK = fmt.Sprintf("wss://%s/", TESTNET_NODE_URL)
var LOCAL_NODE_GETWORK = fmt.Sprintf("ws://%s/", LOCAL_NODE_URL)

var LOCAL_WALLET_URL = fmt.Sprintf("127.0.0.1:%d", WALLET_RPC_PORT)
var LOCAL_WALLET_RPC = fmt.Sprintf("http://%s/json_rpc", LOCAL_WALLET_URL)
var LOCAL_WALLET_WS = fmt.Sprintf("ws://%s/json_rpc", LOCAL_WALLET_URL)

var LOCAL_XSWD_URL = fmt.Sprintf("127.0.0.1:%d", XSWD_PORT)
var LOCAL_XSWD_WS = fmt.Sprintf("ws://%s/xswd", LOCAL_XSWD_URL)

const XELIS_ASSET = `0000000000000000000000000000000000000000000000000000000000000000`
