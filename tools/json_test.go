package tools_test

import (
	"reflect"
	"testing"

	"github.com/anihex/server-utils/tools"
)

func TestStringify(t *testing.T) {
	tt := []struct {
		Name   string
		Param  interface{}
		Result string
	}{
		{
			Name:   "simple number",
			Param:  5,
			Result: "5",
		},
		{
			Name:   "simple string",
			Param:  "Foo.Bar",
			Result: `"Foo.Bar"`,
		},
		{
			Name: "map[string]interface{}",
			Param: map[string]interface{}{
				"int": 5,
				"str": "foo.bar",
			},
			Result: `{"int":5,"str":"foo.bar"}`,
		},
	}
	for _, tc := range tt {
		if result := tools.Stringify(tc.Param); !reflect.DeepEqual(result, tc.Result) {
			t.Errorf("Stringify() = %v, want %v", result, tc.Result)
		}
	}
}
