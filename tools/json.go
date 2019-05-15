package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anihex/json"
)

// BindJSON assumes the body of the request contains a JSON Object and attempts
// to decode it. It will bind the object to the target if, possible.
func BindJSON(r *http.Request, target interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, target)

	return err
}

// Stringify takes an object and returns a JSON string of the object.
func Stringify(i interface{}) string {
	var result string

	b, err := json.Marshal(i)
	if err != nil {
		return ""
	}

	result = fmt.Sprintf("%s", b)

	return result
}
