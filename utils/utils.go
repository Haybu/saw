package utils

import (
	"encoding/json"
	"github.com/jmoiron/jsonq"
	"strings"
)

func JsonValue(body string, parameter string) string {

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	name, err := jq.String(parameter)

	if err != nil {
		msg := "error reading posted parameter " + parameter
		panic(msg)
	}

	return name
}

// check simply panics if there is an error.
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
