package address

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
)

type ValueType int

var Bool ValueType = 0
var String ValueType = 1
var U8 ValueType = 2
var U16 ValueType = 3
var U32 ValueType = 4
var U64 ValueType = 5
var U128 ValueType = 6
var Hash ValueType = 7

var ErrExpectedValue = errors.New("expected a value")
var ErrExpectedArray = errors.New("expected an array")
var ErrExpectedElement = errors.New("expected an element")
var ErrExpectedMap = errors.New("expected a map")

func ErrUnexpectedValueType(valueType ValueType) error {
	return fmt.Errorf("unexpected value type %d", valueType)
}

type DataValue interface{}

type DataElement struct {
	Value  DataValue
	Array  []DataElement
	Fields map[DataValue]DataElement
}

type DataValueReader struct {
	Reader *bytes.Reader
}

func (d *DataValueReader) Read() (dataElement DataElement, err error) {
	dataElementType, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	switch dataElementType {
	case 0: // Value
		var value DataValue
		value, err = d.readDataValue()
		if err != nil {
			return
		}

		dataElement = DataElement{Value: value}
	case 1: // Array
		var size byte
		size, err = d.Reader.ReadByte()
		if err != nil {
			return
		}

		var values []DataElement
		for i := 0; i < int(size); i++ {
			var value DataElement
			value, err = d.Read()
			if err != nil {
				return
			}

			values = append(values, value)
		}

		dataElement = DataElement{Array: values}
	case 2: // Fields / Map
		var size byte
		size, err = d.Reader.ReadByte()
		if err != nil {
			return
		}

		fields := make(map[DataValue]DataElement, 0)
		for i := 0; i < int(size); i++ {
			var key DataValue
			key, err = d.readDataValue()
			if err != nil {
				return
			}

			var value DataElement
			value, err = d.Read()
			if err != nil {
				return
			}

			fields[key] = value
		}

		dataElement = DataElement{Fields: fields}
	default:
		err = fmt.Errorf("invalid data type")
		return
	}

	return
}

func (d *DataValueReader) readValueType() (valueType ValueType, err error) {
	data, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	switch data {
	case 0:
		valueType = Bool
	case 1:
		valueType = String
	case 2:
		valueType = U8
	case 3:
		valueType = U16
	case 4:
		valueType = U32
	case 5:
		valueType = U64
	case 6:
		valueType = U128
	case 7:
		valueType = Hash
	default:
		// error
	}

	return
}

func (d *DataValueReader) readBool() (value bool, err error) {
	data, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	switch data {
	case 0:
		value = false
	case 1:
		value = true
	}

	return
}

func (d *DataValueReader) readString() (value string, err error) {
	size, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	data := make([]byte, size)
	_, err = d.Reader.Read(data)
	if err != nil {
		return
	}

	value = string(data)
	return
}

func (d *DataValueReader) readBytes(size int) (value []byte, err error) {
	data := make([]byte, size)
	_, err = d.Reader.Read(data)
	if err != nil {
		return
	}

	value = data
	return
}

func (d *DataValueReader) readU8() (value uint8, err error) {
	data, err := d.readBytes(1)
	if err != nil {
		return
	}

	value = uint8(data[0])
	return
}

func (d *DataValueReader) readU16() (value uint16, err error) {
	data, err := d.readBytes(2)
	if err != nil {
		return
	}

	value = binary.BigEndian.Uint16(data)
	return
}

func (d *DataValueReader) readU32() (value uint32, err error) {
	data, err := d.readBytes(4)
	if err != nil {
		return
	}

	value = binary.BigEndian.Uint32(data)
	return
}

func (d *DataValueReader) readU64() (value uint64, err error) {
	data, err := d.readBytes(8)
	if err != nil {
		return
	}

	value = binary.BigEndian.Uint64(data)
	return
}

func (d *DataValueReader) readU128() (value big.Int, err error) {
	data, err := d.readBytes(16)
	if err != nil {
		return
	}

	value.SetBytes(data)
	return
}

func (d *DataValueReader) readDataValue() (value DataValue, err error) {
	dataType, err := d.readValueType()
	if err != nil {
		return
	}

	switch dataType {
	case Bool:
		value, err = d.readBool()
	case String:
		value, err = d.readString()
	case U8:
		value, err = d.readU8()
	case U16:
		value, err = d.readU16()
	case U32:
		value, err = d.readU32()
	case U64:
		value, err = d.readU64()
	case U128:
		value, err = d.readU128()
	case Hash:
		value, err = d.readBytes(32)
	}

	return
}
