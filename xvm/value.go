package xvm

// ValueType represents the kind of value stored in ValueCell
type ValueType string

const (
	// All these types are primitive types
	Null   ValueType = "null"
	U8     ValueType = "u8"
	U16    ValueType = "u16"
	U32    ValueType = "u32"
	U64    ValueType = "u64"
	U128   ValueType = "u128"
	U256   ValueType = "u256"
	String ValueType = "string"
	Bool   ValueType = "boolean"
	Range  ValueType = "range"
	// Opaque serde is custom, depending how it was registered
	Opaque ValueType = "opaque"

	PrimitiveKind ValueType = "primitive"

	// value is []ValueCell
	ObjectKind ValueType = "object"
	// value is hex string
	BytesKind ValueType = "bytes"
	// value is [][2]ValueCell (key,value)
	MapKind ValueType = "map"
)

type ValueCell struct {
	Type  ValueType   `json:"type"`
	Value interface{} `json:"value,omitempty"`
}

func (ty *ValueType) IsPrimitive() bool {
	switch *ty {
	case Null, U8, U16, U32, U64, U128, U256, String, Bool, Range, Opaque:
		return true
	default:
		return false
	}
}

func NewPrimitive(ty ValueType, value interface{}) ValueCell {
	if !ty.IsPrimitive() {
		panic("NewPrimitive called with non-primitive type")
	}

	return ValueCell{
		Type: PrimitiveKind,
		Value: ValueCell{
			Type:  ty,
			Value: value,
		},
	}
}

func NewBytes(data string) ValueCell {
	return ValueCell{
		Type:  BytesKind,
		Value: data,
	}
}

func NewObject(values []ValueCell) ValueCell {
	return ValueCell{
		Type:  ObjectKind,
		Value: values,
	}
}

func NewMap(entries [][2]ValueCell) ValueCell {
	return ValueCell{
		Type:  MapKind,
		Value: entries,
	}
}
