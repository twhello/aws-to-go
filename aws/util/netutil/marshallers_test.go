package netutil

import (
	"testing"
	"time"
)

type StructVal struct {
	ValueS string
	ValueI int
	Value  bool
}

type MarshalTest struct {
	StringVal    string `name:"OtherStringVal,omitempty"`
	IntVal       int64  `name:"OtherIntVal"`
	FloatVal     float64
	BoolVal      bool
	TimeVal      time.Time         `name:",omitempty" format:"2006-01-02T15:04:05.999999999Z07:00"`
	TimeValPtr   *time.Time        `name:",omitempty" format:"2006-01-02T15:04:05.999999999Z07:00"`
	StrMap       map[string]string `name:"Str-Map-*"`
	StrArray     []string          `name:"Str.Array.#" base:"2"` // base default = 0
	IntArray     []int64           `name:"Int.Array.#"`
	StructVal    StructVal         `name:"StructVal."`
	StructValPrt *StructVal        `name:"StructValPtr."`
}

func TestMarshalUnmarshal(t *testing.T) {

	tyme := time.Now()

	marTest := &MarshalTest{"String Value", 123, 123.45, true, tyme, &tyme,
		map[string]string{"ValOne": "Value One", "ValTwo": "Value Two"},
		[]string{"A", "B", "C", "D"},
		[]int64{1, 2, 3, 4},
		StructVal{"Value", 123, true},
		&StructVal{"Value2", 456, true},
	}
	t.Logf("MarshalTest:\n%+v\n", marTest)

	vals := MarshalValues(marTest)
	t.Logf("url.Values:\n%+v\n", vals)

	marBack := new(MarshalTest)
	UnmarshalValues(vals, marBack)
	t.Logf("UnmarshalValues:\n%+v %+v\n", marBack, marBack.StructValPrt)

	hdr := MarshalHeader(marTest)
	t.Logf("http.Header:\n%+v\n", hdr)

	marBack = new(MarshalTest)
	UnmarshalHeader(hdr, marBack)
	t.Logf("UnmarshalHeader:\n%+v %+v\n", marBack, marBack.StructValPrt)
}
