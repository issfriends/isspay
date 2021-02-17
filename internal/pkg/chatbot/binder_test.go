package chatbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	FieldStr      string   `json:"fieldStr"`
	FieldInt      int64    `json:"fieldInt"`
	FieldUInt     uint64   `json:"fieldUInt"`
	FieldPtrInt   *int64   `json:"fieldPtrInt"`
	FieldFloat    float32  `json:"fieldFloat"`
	FieldSliceInt []int    `json:"fieldSliceInt"`
	FieldSliceStr []string `json:"fieldSliceStr"`
	FieldBool     bool     `json:"fieldBool"`
	NotBind       string
}

func TestBind(t *testing.T) {
	data := map[string]string{
		"fieldStr":      "testStr",
		"fieldInt":      "-1",
		"fieldUInt":     "12345",
		"fieldPtrInt":   "100",
		"fieldFloat":    "0.0023",
		"fieldSliceInt": "1,2,3,4,5",
		"fieldSliceStr": "a,b,c,d,e,fg",
		"fieldBool":     "true",
	}

	ts := &testStruct{}
	bind(data, ts, "json")
	assert.Equal(t, ts.FieldStr, "testStr")
	assert.Equal(t, ts.FieldInt, int64(-1))
	assert.Equal(t, ts.FieldUInt, uint64(12345))
	assert.NotNil(t, ts.FieldPtrInt)
	assert.Equal(t, *ts.FieldPtrInt, int64(100))
	assert.Equal(t, ts.FieldFloat, float32(0.0023))
	assert.True(t, ts.FieldBool)
	assert.Equal(t, ts.FieldSliceInt, []int{1, 2, 3, 4, 5})
	assert.Equal(t, ts.FieldSliceStr, []string{"a", "b", "c", "d", "e", "fg"})
	assert.Empty(t, ts.NotBind)
}
