package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncode(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)
	customer := Customer{
		FirstName:  "Ahmad",
		MiddleName: "Najibul",
		LastName:   "Subairi",
	}
	encoder.Encode(customer)
}
