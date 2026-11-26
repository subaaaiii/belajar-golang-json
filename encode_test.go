package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJSON(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
func TestEncode(t *testing.T) {
	LogJSON("subairi")
	LogJSON(277)
	LogJSON(false)
	LogJSON([]string{"subairi", "joki", "nopal"})
}
