package address

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"
)

type ValueType int

var BoolType ValueType = 0
var StringType ValueType = 1
var U8Type ValueType = 2
var U16Type ValueType = 3
var U32Type ValueType = 4
var U64Type ValueType = 5
var U128Type ValueType = 6
var HashType ValueType = 7

type Hash [32]byte

type DataElementType int

var DataElementValue DataElementType = 0
var DataElementArray DataElementType = 1
var DataElementFields DataElementType = 2

var ErrMaxStringSize = errors.New("string max limit is 255 bytes")

func ErrUnsupportedValue(value DataValue) error {
	return fmt.Errorf("unsupported value type %v", value)
}

type DataValue interface{}

type DataElement struct {
	Value  DataValue
	Array  []DataElement
	Fields map[DataValue]DataElement
}

// use this function to convert in a valid map for json.Marshal()
func (d DataElement) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	if d.Value != nil {
		switch value := d.Value.(type) {
		case big.Int:
			result["value"] = value.String()
		default:
			result["value"] = value
		}
	}

	if d.Array != nil {
		var array []interface{}
		for _, item := range d.Array {
			array = append(array, item.ToMap())
		}

		result["array"] = array
	}

	if d.Fields != nil {
		fields := make(map[string]interface{})
		for key, item := range d.Fields {
			sKey := fmt.Sprintf("%v", key)
			fields[sKey] = item.ToMap()
		}

		result["fields"] = fields
	}

	return result
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
	case byte(DataElementValue): // Value
		var value DataValue
		value, err = d.readDataValue()
		if err != nil {
			return
		}

		dataElement = DataElement{Value: value}
	case byte(DataElementArray): // Array
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
	case byte(DataElementFields): // Fields / Map
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

	valueType = ValueType(data)
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
	data, err := d.Reader.ReadByte()
	if err != nil {
		return
	}

	value = uint8(data)
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

func (d *DataValueReader) readHash() (value Hash, err error) {
	data, err := d.readBytes(32)
	if err != nil {
		return
	}

	copy(value[:], data)
	return
}

func (d *DataValueReader) readDataValue() (value DataValue, err error) {
	valueType, err := d.readValueType()
	if err != nil {
		return
	}

	switch valueType {
	case BoolType:
		value, err = d.readBool()
	case StringType:
		value, err = d.readString()
	case U8Type:
		value, err = d.readU8()
	case U16Type:
		value, err = d.readU16()
	case U32Type:
		value, err = d.readU32()
	case U64Type:
		value, err = d.readU64()
	case U128Type:
		value, err = d.readU128()
	case HashType:
		value, err = d.readHash()
	}

	return
}

type DataValueWriter struct {
	Writer io.Writer
}

func (d *DataValueWriter) Write(dataElement DataElement) (err error) {
	if dataElement.Value != nil {
		err = d.writeByte(byte(DataElementValue))
		if err != nil {
			return
		}

		err = d.writeValue(dataElement.Value)
		if err != nil {
			return
		}
	}

	if dataElement.Array != nil {
		err = d.writeByte(byte(DataElementArray))
		if err != nil {
			return
		}

		err = d.writeByte(byte(len(dataElement.Array)))
		if err != nil {
			return
		}

		for _, item := range dataElement.Array {
			err = d.Write(item)
			if err != nil {
				return
			}
		}
	}

	if dataElement.Fields != nil {
		err = d.writeByte(byte(DataElementFields))
		if err != nil {
			return
		}

		err = d.writeByte(byte(len(dataElement.Fields)))
		if err != nil {
			return
		}

		for key, item := range dataElement.Fields {
			err = d.writeValue(key)
			if err != nil {
				return
			}

			err = d.Write(item)
			if err != nil {
				return
			}
		}
	}

	return
}

func (d *DataValueWriter) writeByte(value byte) (err error) {
	data := make([]byte, 1)
	data[0] = value

	_, err = d.Writer.Write(data)
	if err != nil {
		return
	}

	return
}

func (d *DataValueWriter) writeBool(value bool) (err error) {
	data := make([]byte, 1)

	if value {
		data[0] = 1
	} else {
		data[0] = 0
	}

	_, err = d.Writer.Write(data)
	if err != nil {
		return
	}

	return
}

func (d *DataValueWriter) writeString(value string) (err error) {
	buf := bytes.NewBufferString(value)

	if buf.Len() > 255 {
		err = ErrMaxStringSize
		return
	}

	size := byte(buf.Len())
	err = d.writeByte(size)
	if err != nil {
		return
	}

	_, err = d.Writer.Write(buf.Bytes())
	if err != nil {
		return
	}

	return
}

func (d *DataValueWriter) writeU16(value uint16) (err error) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, value)
	_, err = d.Writer.Write(data)
	return
}

func (d *DataValueWriter) writeU32(value uint32) (err error) {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, value)
	_, err = d.Writer.Write(data)
	return
}

func (d *DataValueWriter) writeU64(value uint64) (err error) {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, value)
	_, err = d.Writer.Write(data)
	return
}

func (d *DataValueWriter) writeU128(value big.Int) (err error) {
	_, err = d.Writer.Write(value.Bytes())
	return
}

func (d *DataValueWriter) writeValue(value DataValue) (err error) {
	switch value := value.(type) {
	case bool:
		err = d.writeByte(byte(BoolType))
		if err != nil {
			return
		}

		err = d.writeBool(value)
	case string:
		err = d.writeByte(byte(StringType))
		if err != nil {
			return
		}

		err = d.writeString(value)
	case uint8:
		err = d.writeByte(byte(U8Type))
		if err != nil {
			return
		}

		err = d.writeByte(byte(value))
	case uint16:
		err = d.writeByte(byte(U16Type))
		if err != nil {
			return
		}

		err = d.writeU16(value)
	case uint32:
		err = d.writeByte(byte(U32Type))
		if err != nil {
			return
		}

		err = d.writeU32(value)
	case uint64:
		err = d.writeByte(byte(U64Type))
		if err != nil {
			return
		}

		err = d.writeU64(value)
	case big.Int:
		err = d.writeByte(byte(U128Type))
		if err != nil {
			return
		}

		err = d.writeU128(value)
	case Hash:
		err = d.writeByte(byte(HashType))
		if err != nil {
			return
		}

		_, err = d.Writer.Write(value[:])
	default:
		err = ErrUnsupportedValue(value)
		return
	}

	return
}
