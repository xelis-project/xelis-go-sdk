package address

import (
	"fmt"
	"testing"
)

var MAINNET_ADDR = "xel:ys4peuzztwl67rzhsdu0yxfzwcfmgt85uu53hycpeeary7n8qvysqmxznt0"

func TestAddressFromString(t *testing.T) {
	address, err := NewAddressFromString(MAINNET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", address)

	addr, err := address.Format()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)

	if addr != MAINNET_ADDR {
		t.Fail()
		t.Logf("Expected %s, got %s", MAINNET_ADDR, addr)
	}
}

func TestIntegratedAddress(t *testing.T) {
	address, err := NewAddressFromString("xet:6eadzwf5xdacts6fs4y3csmnsmy4mcxewqt3xyygwfx0hm0tm32szqsrqyzkjar9d4esyqgpq4ehwmmjvsqqypgpq45x2mrvduqqzpthdaexceqpq4mk7unywvqsgqqpq4yx2mrvduqqzp2hdaexceqqqyzxvun0d5qqzp2cg4xyj5ct5udlg")
	if err != nil {
		t.Fatal(err)
	}

	if !address.Integrated {
		t.Fail()
		t.Logf("Expected address to be integrated")
	}

	t.Logf("%+v\n", address.ExtraData)

	if address.ExtraData.Value != nil {
		fmt.Println(address.ExtraData.Value)
	}

	if address.ExtraData.Array != nil {
		for _, value := range address.ExtraData.Array {
			fmt.Println(value)
		}
	}

	if address.ExtraData.Fields != nil {
		for _, value := range address.ExtraData.Fields {
			fmt.Println(value)
		}
	}
}
