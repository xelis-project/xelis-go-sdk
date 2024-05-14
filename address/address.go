package address

import (
	"bytes"
	"fmt"
)

var PrefixAddress string = "xel"
var TestnetPrefixAddress string = "xet"

var ExtraDataLimit = 1024

type Address struct {
	Mainnet    bool
	Integrated bool
	ExtraData  DataElement
	Key        []byte
}

func NewAddress(data []byte, hrp string) (addr *Address, err error) {
	reader := bytes.NewReader(data)

	publicKey := make([]byte, 32)
	_, err = reader.Read(publicKey)
	if err != nil {
		return
	}

	addrType, err := reader.ReadByte()
	if err != nil {
		return
	}

	integrated := false
	var dataElement DataElement

	switch addrType {
	case 0:
		// do nothing
	case 1:
		integrated = true

		dataValueReader := &DataValueReader{Reader: reader}
		dataElement, err = dataValueReader.Read()
		if err != nil {
			return
		}

		if reader.Size() > int64(ExtraDataLimit) {
			// error
			return
		}
	default:
		err = fmt.Errorf("invalid address type")
		return
	}

	addr = &Address{
		Mainnet:    hrp == PrefixAddress,
		Key:        publicKey,
		Integrated: integrated,
		ExtraData:  dataElement,
	}

	return
}

func NewAddressFromString(address string) (addr *Address, err error) {
	hrp, decoded, err := decode(address)
	if err != nil {
		return
	}

	if hrp != PrefixAddress && hrp != TestnetPrefixAddress {
		return
	}

	bits, err := convertBits(decoded, 5, 8, false)
	if err != nil {
		return
	}

	addr, err = NewAddress(bits, hrp)
	if err != nil {
		return
	}

	return
}

func (a *Address) Format() (addr string, err error) {
	var buf bytes.Buffer
	buf.Write(a.Key)
	if a.Integrated {
		buf.Write([]byte{1})
	} else {
		buf.Write([]byte{0})
	}
	data := buf.Bytes()

	bits, err := convertBits(data, 8, 5, true)
	if err != nil {
		return
	}

	hrp := PrefixAddress
	if !a.Mainnet {
		hrp = TestnetPrefixAddress
	}

	addr, err = encode(hrp, bits)
	return
}
