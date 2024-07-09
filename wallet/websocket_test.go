package wallet

import (
	"log"
	"sync"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func useWSLocal(t *testing.T) (wallet *WebSocket) {
	wallet, err := NewWebSocket(config.LOCAL_WALLET_WS, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestWSGetVersion(t *testing.T) {
	wallet := useWSLocal(t)
	version, err := wallet.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", version)
	wallet.Close()
}

func TestWSGetNetwork(t *testing.T) {
	wallet := useWSLocal(t)
	network, err := wallet.GetNetwork()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", network)
	wallet.Close()
}

func TestWSNewTopoheight(t *testing.T) {
	wallet := useWSLocal(t)

	var wg sync.WaitGroup
	wg.Add(1)
	err := wallet.NewTopoheightFunc(func(newTopoheight uint64, err error) {
		if err != nil {
			log.Fatal(err)
		}

		t.Log(newTopoheight)
		wg.Done()
	})
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	wallet.Close()
}

func TestWSNewTopoheightChannel(t *testing.T) {
	wallet := useWSLocal(t)

	newTopoheight, newTopoheightErr, err := wallet.NewTopoheightChannel()
	if err != nil {
		t.Fatal(err)
	}

	select {
	case topoheight := <-newTopoheight:
		t.Log(topoheight)
	case err := <-newTopoheightErr:
		t.Fatal(err)
	}

	close(newTopoheight)
	close(newTopoheightErr)
	wallet.Close()
}

func TestWSOnlineOffline(t *testing.T) {
	wallet := useWSLocal(t)

	var wg sync.WaitGroup
	wg.Add(2)

	err := wallet.OnlineFunc(func() {
		t.Log("Online")
		wg.Done()
	})
	if err != nil {
		log.Fatal(err)
	}

	err = wallet.OfflineFunc(func() {
		t.Log("Offline")
		wg.Done()
	})
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	wallet.Close()
}

func TestConnectionErr(t *testing.T) {
	wallet := useWSLocal(t)

	err := <-wallet.ConnectionErr()
	t.Log(err)
}
