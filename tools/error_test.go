package tools_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anihex/server-utils/tools"
)

func TestErrIfTrue(t *testing.T) {
	tt := []struct {
		Name      string
		Condition bool
		Message   string
		Params    []interface{}
		Result    error
	}{
		{
			Name:      "is true",
			Condition: true,
			Message:   "this condition is %v",
			Params:    []interface{}{true},
			Result:    fmt.Errorf("this condition is true"),
		},
		{
			Name:      "is false",
			Condition: false,
			Message:   "this condition is %v",
			Params:    []interface{}{false},
			Result:    nil,
		},
	}

	for _, tc := range tt {
		result := tools.ErrIfTrue(tc.Condition, tc.Message, tc.Params...)
		if !reflect.DeepEqual(result, tc.Result) {
			t.Errorf("case %s failed. '%v' expected, got '%v'", tc.Name, tc.Result, result)
		}
	}
}
