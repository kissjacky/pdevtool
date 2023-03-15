package cmd

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRemoveNulls(t *testing.T) {
	a := []byte(`{"a":1,"b":null,"c":{"d":null}}`)
	var j interface{}
	json.Unmarshal(a, &j)
	fmt.Println(j)
	removeNulls(j)
	fmt.Println(j)
}
