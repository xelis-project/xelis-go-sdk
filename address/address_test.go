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

func TestInvalidAddress(t *testing.T) {
	address, err := NewAddressFromString("xel:ys4peuzztwl67rzhsdu0yxfzwcfmgt85uu53hycpeeary7n8qvysqmxznt1")
	if err != nil {
		t.Log(address, err)
	}

	valid, err := IsValidAddress(MAINNET_ADDR)
	if !valid {
		t.Log(err)
		t.Fail()
	}
}

func TestIntegratedAddress(t *testing.T) {
	address, err := NewAddressFromString("xet:6eadzwf5xdacts6fs4y3csmnsmy4mcxewqt3xyygwfx0hm0tm32szqsrqyzkjar9d4esyqgpq4ehwmmjvsqqypgpq45x2mrvduqqzpthdaexceqpq4mk7unywvqsgqqpq4yx2mrvduqqzp2hdaexceqqqyzxvun0d5qqzp2cg4xyj5ct5udlg")
	if err != nil {
		t.Fatal(err)
	}

	if !address.IsIntegrated() {
		t.Fail()
		t.Logf("Expected address to be integrated")
	}

	extraData := address.GetExtraData()
	t.Logf("%+v\n", extraData)

	if extraData.Value != nil {
		fmt.Println(extraData.Value)
	}

	if extraData.Array != nil {
		for _, value := range extraData.Array {
			fmt.Println(value)
		}
	}

	if extraData.Fields != nil {
		for _, value := range extraData.Fields {
			fmt.Println(value)
		}
	}
}

func TestAddressExtraDataValue(t *testing.T) {
	address, err := NewAddressFromString(MAINNET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	address.SetExtraData(&DataElement{Value: true})

	addr, err := address.Format()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)

	addressCopy, err := NewAddressFromString(addr)
	if err != nil {
		return
	}

	extraData := addressCopy.GetExtraData()
	t.Logf("%+v\n", extraData)
}

func TestAddressExtraDataArray(t *testing.T) {
	address, err := NewAddressFromString(MAINNET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	address.SetExtraData(&DataElement{
		Array: []DataElement{
			{Value: true},
			{Value: "test"},
		},
	})

	addr, err := address.Format()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)

	addressCopy, err := NewAddressFromString(addr)
	if err != nil {
		return
	}

	extraData := addressCopy.GetExtraData()
	t.Logf("%+v\n", extraData)
}

func TestAddressExtraDataFields(t *testing.T) {
	address, err := NewAddressFromString(MAINNET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	fields := make(map[DataValue]DataElement, 0)
	fields["hello"] = DataElement{Value: "world"}

	address.SetExtraData(&DataElement{
		Fields: fields,
	})

	addr, err := address.Format()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)

	addressCopy, err := NewAddressFromString(addr)
	if err != nil {
		return
	}

	extraData := addressCopy.GetExtraData()
	t.Logf("%+v\n", extraData)
}

func TestAddressClearExtraData(t *testing.T) {
	address, err := NewAddressFromString(MAINNET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", address)

	address.SetExtraData(&DataElement{Value: "test"})

	addr, err := address.Format()
	if err != nil {
		t.Fatal(err)
	}

	address2, err := NewAddressFromString(addr)
	if err != nil {
		t.Fatal(err)
	}

	address2.ClearExtraData()
	addr2, err := address2.Format()
	if err != nil {
		t.Fatal(err)
	}

	if addr2 != MAINNET_ADDR {
		t.Fail()
		t.Logf("Expected %s, got %s", MAINNET_ADDR, addr)
	}
}
