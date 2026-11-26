package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Ahmad",
		MiddleName: "Muh",
		LastName:   "Subairi",
		Age:        18,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
