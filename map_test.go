package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	product := map[string]any{
		"id":    "p-0001",
		"name":  "Apple macbook pro",
		"price": 200000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"p-0001","name":"Apple macbook pro","price":200000}`
	jsonByte := []byte(jsonString)
	var result map[string]interface{}
	err := json.Unmarshal(jsonByte, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["name"])
}
