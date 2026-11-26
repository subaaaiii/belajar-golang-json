package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Ahmad",
		MiddleName: "Najibul",
		LastName:   "Subairi",
		Hobbies:    []string{"Gaming, Reading, Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeArray(t *testing.T) {
	jsonString := `{"FirstName":"Ahmad","MiddleName":"Najibul","LastName":"Subairi","Age":0,"Married":false,"Hobbies":["Gaming, Reading, Coding"]}`
	jsonByte := []byte(jsonString)
	cust := &Customer{}

	json.Unmarshal(jsonByte, cust)

	fmt.Println(cust)
	fmt.Println(cust.FirstName)
	fmt.Println(cust.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Ahmad",
		MiddleName: "Najibul",
		LastName:   "Subairi",
		Hobbies:    []string{"Gaming, Reading, Coding"},
		Addresses: []Address{
			{
				Street:     "Jalan lagi rusan",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan belum ada",
				Country:    "Malaysia",
				PostalCode: "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeArrayComplex(t *testing.T) {
	jsonString := `{"FirstName":"Ahmad","MiddleName":"Najibul","LastName":"Subairi","Age":0,"Married":false,"Hobbies":["Gaming, Reading, Coding"],"Addresses":[{"Street":"Jalan lagi rusan","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan belum ada","Country":"Malaysia","PostalCode":"8888"}]}`
	jsonByte := []byte(jsonString)
	cust := &Customer{}

	json.Unmarshal(jsonByte, cust)

	fmt.Println(cust)
	fmt.Println(cust.FirstName)
	fmt.Println(cust.Hobbies)
	fmt.Println(cust.Addresses)
	fmt.Println(cust.Addresses[0].Street)
}

func TestJsonArrayOnly(t *testing.T) {

	Addresses := []Address{
		{
			Street:     "Jalan lagi rusan",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan belum ada",
			Country:    "Malaysia",
			PostalCode: "8888",
		},
	}

	bytes, _ := json.Marshal(Addresses)
	fmt.Println(string(bytes))
}
func TestDecodeArrayOnly(t *testing.T) {
	jsonString := `[{"Street":"Jalan lagi rusan","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan belum ada","Country":"Malaysia","PostalCode":"8888"}]`
	jsonByte := []byte(jsonString)
	addresses := &[]Address{}

	json.Unmarshal(jsonByte, addresses)

	fmt.Println(addresses)

}
