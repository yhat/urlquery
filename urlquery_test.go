package urlquery

import (
	"testing"
)

type marshalTest struct {
	S interface{}
	V map[string]string
}

var marshalTests = []marshalTest{
	marshalTest{
		S: struct {
			V1        string `url:"v4"`
			V2        string `url:"v3"`
			lowerName string
		}{"hi", "bye", "alongstring"},
		V: map[string]string{"v4": "hi", "v3": "bye", "lowerName": ""},
	},
	marshalTest{
		S: struct {
			V1 string
			V2 int `url:"myval"`
			V3 int
		}{"v1", 232, -3},
		V: map[string]string{"V1": "v1", "myval": "232", "V3": "-3"},
	},
}

func TestMarshal(t *testing.T) {
	for _, test := range marshalTests {
		result := Marshal(test.S)
		for k, v := range test.V {
			if result.Get(k) != v {
				t.Errorf("PBV: Expected '%s' got '%s'", v, result.Get(k))
			}
		}
		result = Marshal(&test.S)
		for k, v := range test.V {
			if result.Get(k) != v {
				t.Errorf("PBP: Expected '%s' got '%s'", v, result.Get(k))
			}
		}
	}
}

func TestUnmarshal(t *testing.T) {
}
