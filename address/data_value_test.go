package address

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"math/big"
	"testing"
)

func writeDataElement(dataElement DataElement) (data []byte, err error) {
	var buf bytes.Buffer
	dataValueWriter := DataValueWriter{Writer: &buf}
	err = dataValueWriter.Write(dataElement)
	if err != nil {
		return
	}

	data = buf.Bytes()
	return
}

func readDataElement(data []byte) (dataElement DataElement, err error) {
	dataValueReader := DataValueReader{Reader: bytes.NewReader(data)}
	dataElement, err = dataValueReader.Read()
	return
}

func TestDataElementBoolValue(t *testing.T) {
	dataElement := DataElement{Value: true}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %t, got %t", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementStringValue(t *testing.T) {
	dataElement := DataElement{Value: "test"}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %s, got %s", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU8Value(t *testing.T) {
	dataElement := DataElement{Value: uint8(122)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU16Value(t *testing.T) {
	dataElement := DataElement{Value: uint16(34566)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU32Value(t *testing.T) {
	dataElement := DataElement{Value: uint32(6767456)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU64Value(t *testing.T) {
	dataElement := DataElement{Value: uint64(345765875678)}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	if dataElement.Value != dataElementCopy.Value {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementU128Value(t *testing.T) {
	var nbr big.Int
	nbr.SetString("35467456745674956794567", 10)

	dataElement := DataElement{Value: nbr}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	v1 := dataElement.Value.(big.Int)
	v2 := dataElement.Value.(big.Int)

	if v1.Cmp(&v2) != 0 {
		t.Fail()
		t.Logf("Expected %d, got %d", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementHashValue(t *testing.T) {
	var hash Hash

	shaHash := sha256.Sum256([]byte("asd"))
	copy(hash[:], shaHash[:])

	dataElement := DataElement{Value: hash}
	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	v1 := dataElement.Value.(Hash)
	v2 := dataElementCopy.Value.(Hash)

	if !bytes.Equal(v1[:], v2[:]) {
		t.Fail()
		t.Logf("Expected %s, got %s", dataElement.Value, dataElementCopy.Value)
	}
}

func TestDataElementArray(t *testing.T) {
	array := []DataElement{
		{Value: true},
		{Value: "test"},
	}

	dataElement := DataElement{Array: array}

	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	for i, item := range array {
		if dataElementCopy.Array[i].Value != item.Value {
			t.Fail()
			t.Logf("Expected %v, got %v", item.Value, dataElementCopy.Array[i].Value)
		}
	}
}

func TestDataElementFields(t *testing.T) {
	fields := make(map[DataValue]DataElement, 0)
	fields["hello"] = DataElement{Value: "world"}

	dataElement := DataElement{Fields: fields}

	data, err := writeDataElement(dataElement)
	if err != nil {
		t.Fatal(err)
	}

	dataElementCopy, err := readDataElement(data)
	if err != nil {
		t.Fatal(err)
	}

	for key, element := range fields {
		if dataElementCopy.Fields[key].Value != element.Value {
			t.Fail()
			t.Logf("Expected %v, got %v", element.Value, dataElementCopy.Fields[key].Value)
		}
	}
}

func TestLongStringMaxLimit(t *testing.T) {
	// max 255 bytes for string
	_, err := writeDataElement(DataElement{Value: "woenrbowirentboiejwrntbpoijewnrtbpenrptbjnepritjbnperijtnbpijewnrtbpjnerptbnjperkjtbnperkjtnbpsdfgsergwngio453gn45oign345iogjnwosiwejrngwpo34i5ny3[45oyhi3n4p5[hokn3p4o5nhekjrntbpkjewnrtpbkjnwerptkbjnpwkrjntbperkjntbpkwerntpbjkenrptbkjnwpekjnrwkpenrfpbknweprkbjnwperkbjn"})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestToMap(t *testing.T) {
	var hash Hash

	shaHash := sha256.Sum256([]byte("asd"))
	copy(hash[:], shaHash[:])

	var bigNumber big.Int
	bigNumber.SetString("2093458230498572039452039485702938475", 10)

	dataElement := DataElement{
		Value: bigNumber,
		Array: []DataElement{
			{Value: true},
			{Value: "test"},
			{Value: uint8(34)},
			{Value: uint16(34523)},
			{Value: uint32(3452305469)},
			{Value: uint64(3452305469567567456)},
			{
				Array: []DataElement{
					{Value: "sub_array"},
				},
			},
		},
		Fields: map[DataValue]DataElement{
			34:             {Value: false},
			"hello":        {Value: "world"},
			23456923846034: {Value: hash},
		},
	}

	sMap := dataElement.ToMap()

	jsonBytes, err := json.Marshal(sMap)
	if err != nil {
		t.Fatal(err)
	}

	jsonString := string(jsonBytes)

	t.Logf("%+v", sMap)
	t.Logf("%+v", string(jsonString))
}
