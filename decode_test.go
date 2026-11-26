package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonString := `{"FirstName":"Ahmad","MiddleName":"Muh","LastName":"Subairi","Age":18,"Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)

}
