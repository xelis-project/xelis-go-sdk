package xvm

import (
	"encoding/json"
	"math/big"
	"testing"
)

func TestConstantOptional(t *testing.T) {
	data, err := json.MarshalIndent(NewPrimitive(String, nil), "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}

func TestConstantArray(t *testing.T) {
	u128 := new(big.Int)
	u128.SetString("340282366920938463463374607431768211455", 10)

	u256 := new(big.Int)
	u256.SetString("115792089237316195423570985008687907853269984665640564039457584007913129659255", 10)

	data, err := json.MarshalIndent([]ValueCell{
		NewPrimitive(Null, nil),
		NewPrimitive(String, "Hello, World!"),
		NewPrimitive(Bool, true),
		NewBytes("48656c6c6f2c20576f726c6421"), // "Hello, World!" in hex
		NewPrimitive(U8, uint8(255)),
		NewPrimitive(U16, uint16(65535)),
		NewPrimitive(U32, uint32(4294967295)),
		NewPrimitive(U64, uint64(18446744073709551615)),
		NewPrimitive(U128, u128),
		NewPrimitive(U256, u256),
	}, "", "  ")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}

func TestValueRange(t *testing.T) {
	data, err := json.MarshalIndent(NewPrimitive(
		Range,
		[]ValueCell{
			{
				Type:  U16,
				Value: uint16(0),
			},
			{
				Type:  U16,
				Value: uint16(50),
			},
		}), "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}
