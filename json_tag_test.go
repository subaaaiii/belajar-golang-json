package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "p-1",
		Name:     "apple macbook pro",
		ImageURL: "https://example.com/image.png",
	}
	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"p-1","name":"apple macbook pro","image_url":"https://example.com/image.png"}`
	jsonByte := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonByte, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)

}
