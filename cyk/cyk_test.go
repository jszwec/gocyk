package cyk

import (
	"reflect"
	"testing"
)

func TestCyk(t *testing.T) {
	data := []struct {
		g map[string][]string
		e CykTable
		w string
	}{
		{
			g: map[string][]string{
				"S": []string{"AB", "BC"},
				"A": []string{"BA", "a"},
				"B": []string{"CC", "b"},
				"C": []string{"AB", "a"},
			},
			e: CykTable{
				[]CykCell{{"b": {}}, {"a": {}}, {"a": {}}, {"b": {}}, {"a": {}}},
				[]CykCell{{"B": {}}, {"A": {}, "C": {}}, {"A": {}, "C": {}}, {"B": {}}, {"A": {}, "C": {}}},
				[]CykCell{{"A": {}, "S": {}}, {"B": {}}, {"C": {}, "S": {}}, {"A": {}, "S": {}}, nil},
				[]CykCell{{}, {"B": {}}, {"B": {}}, nil, nil},
				[]CykCell{{}, {"S": {}, "A": {}, "C": {}}, nil, nil, nil},
				[]CykCell{{"S": {}, "A": {}, "C": {}}, nil, nil, nil, nil},
			},

			w: "baaba",
		},
		{
			g: map[string][]string{
				"S": []string{"AB"},
				"A": []string{"a"},
				"B": []string{"b"},
			},
			e: CykTable{
				[]CykCell{{"a": {}}, {"b": {}}},
				[]CykCell{{"A": {}}, {"B": {}}},
				[]CykCell{{"S": {}}, nil},
			},
			w: "ab",
		},
		{
			g: map[string][]string{
				"S": []string{"a"},
			},
			e: CykTable{
				[]CykCell{{"a": {}}},
				[]CykCell{{"S": {}}},
			},
			w: "a",
		},
	}
	for _, val := range data {
		ct := Cyk(val.g, val.w)
		if !reflect.DeepEqual(val.e, ct) {
			t.Errorf("expected '%#v', got '%#v'", val.e, ct)
		}
	}
}
