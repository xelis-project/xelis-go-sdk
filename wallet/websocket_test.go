package wallet

import (
	"log"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func setupWebSocket(t *testing.T) (wallet *WebSocket) {
	wallet, err := NewWebSocket(config.LOCAL_WALLET_WS, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestWSGetInfo(t *testing.T) {
	wallet := setupWebSocket(t)
	version, err := wallet.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", version)

	network, err := wallet.GetNetwork()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", network)

	wallet.Close()
}
