package dynamodb

import (
	"fmt"
	"strconv"
)

// Represents the data for an attribute. You can set one, and only one, of the elements.
//
// Each attribute in an item is a name-value pair. An attribute can be single-valued or
// multi-valued set. For example, a book item can have title and authors attributes.
// Each book has one title but can have many authors. The multi-valued attribute is a set;
// duplicate values are not allowed.
// [http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeValue.html]
type AttributeValue struct {
	B  []byte   `json:"B,omitempty"`
	BS [][]byte `json:"BS,omitempty"`
	N  string   `json:"N,omitempty"`
	NS []string `json:"NS,omitempty"`
	S  string   `json:"S,omitempty"`
	SS []string `json:"SS,omitempty"`
}

// Returns the data type of the Attribute
func (v AttributeValue) Type() AttributeType {
	switch {
	case len(v.S)+len(v.SS) > 0:
		return STRING
	case len(v.N)+len(v.NS) > 0:
		return NUMBER
	}
	return BINARY
}

// Checks if the Attribute contains any values.
func (v AttributeValue) IsEmpty() bool {
	return len(v.S)+len(v.SS)+len(v.N)+len(v.NS)+len(v.B)+len(v.BS) == 0
}

// A Binary data type (0 to 255).
func (v AttributeValue) Binary() (val []byte) {
	return v.B
}

// A Binary set data type (0 to 255).
func (v AttributeValue) BinarySet() (val [][]byte) {
	return v.BS
}

// A Number data type as int64 (-9223372036854775808 to 9223372036854775807).
func (v AttributeValue) Int() (val int64) {
	val, _ = strconv.ParseInt(v.N, 10, 64)
	return val
}

// A Number set data type as int64 array (-9223372036854775808 to 9223372036854775807).
func (v AttributeValue) IntSet() (val []int64) {
	val = make([]int64, len(v.NS))
	for i, value := range v.NS {
		i64, _ := strconv.ParseInt(value, 10, 64)
		val[i] = i64
	}
	return
}

// A Number data type as uint64 (0 to 18446744073709551615).
func (v AttributeValue) Uint() (val uint64) {
	val, _ = strconv.ParseUint(v.N, 10, 64)
	return val
}

// A Number set data type as uint64 array (0 to 18446744073709551615).
func (v AttributeValue) UintSet() (val []uint64) {
	val = make([]uint64, len(v.NS))
	for i, value := range v.NS {
		i64, _ := strconv.ParseUint(value, 10, 64)
		val[i] = i64
	}
	return
}

// A Number data type as float64.
func (v AttributeValue) Float() (val float64) {
	val, _ = strconv.ParseFloat(v.N, 64)
	return val
}

// A Number set data type as float64 array.
func (v AttributeValue) FloatSet() (val []float64) {
	val = make([]float64, len(v.NS))
	for i, value := range v.NS {
		f64, _ := strconv.ParseFloat(value, 64)
		val[i] = f64
	}
	return
}

// A String data type.
func (v AttributeValue) Value() (val string) {
	return v.S
}

// A String set data type.
func (v AttributeValue) ValueSet() (val []string) {
	return v.SS
}

// String representation of the AttributeValue.
func (v AttributeValue) String() string {
	return fmt.Sprintf("AttributeValue[B:%s, BS:%s, N:%s, NS:%s, S:%s, SS:%s]", v.B, v.BS, v.N, v.NS, v.S, v.SS)
}

/******************************************************************************
 * AttributeValue Initializers
 */

// Creates an AttributeValue of a Binary data type.
func NewBinaryAttributeValue(val []byte) AttributeValue {
	return AttributeValue{B: val}
}

// Creates an AttributeValue of a Binary set data type.
func NewBinaryAttributeSet(val [][]byte) AttributeValue {
	return AttributeValue{BS: val}
}

// Creates an AttributeValue of a Number data type using float64.
func NewFloatAttributeValue(val float64) AttributeValue {
	return AttributeValue{N: strconv.FormatFloat(val, 'f', -1, 64)}
}

// Creates an AttributeValue of a Number set data type using float64 array.
func NewFloatAttributeSet(val []float64) AttributeValue {
	strSet := make([]string, len(val))
	for i, v := range val {
		strSet[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return AttributeValue{NS: strSet}
}

// Creates an AttributeValue of a Number data type using int64.
func NewIntAttributeValue(val int64) AttributeValue {
	return AttributeValue{N: strconv.FormatInt(val, 10)}
}

// Creates an AttributeValue of a Number set data type using int64 array.
func NewIntAttributeSet(val []int64) AttributeValue {
	strSet := make([]string, len(val))
	for i, v := range val {
		strSet[i] = strconv.FormatInt(v, 10)
	}
	return AttributeValue{NS: strSet}
}

// Creates an AttributeValue of a Number data type using uint64.
// (0 to 18446744073709551615)
func NewUintAttributeValue(val uint64) AttributeValue {
	return AttributeValue{N: strconv.FormatUint(val, 10)}
}

// Creates an AttributeValue of a Number set data type using uint64 array.
// (0 to 18446744073709551615)
func NewUintAttributeSet(val []uint64) AttributeValue {
	strSet := make([]string, len(val))
	for i, v := range val {
		strSet[i] = strconv.FormatUint(v, 10)
	}
	return AttributeValue{NS: strSet}
}

// Creates an AttributeValue of a String data type.
func NewAttributeValue(val string) AttributeValue {
	return AttributeValue{S: val}
}

// Creates an AttributeValue of a String set data type.
func NewAttributeSet(val []string) AttributeValue {
	return AttributeValue{SS: val}
}
