package wallet

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

func setupRPC(t *testing.T) (wallet *RPC, ctx context.Context) {
	ctx = context.Background()
	wallet, err := NewRPC(ctx, config.LOCAL_WALLET_RPC, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCGetInfo(t *testing.T) {
	wallet, _ := setupRPC(t)

	version, err := wallet.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)
}
