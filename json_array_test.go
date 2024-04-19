package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArrayJson(t *testing.T) {
	customer := Customer{
		FirstName:  "Bagas",
		MiddleName: "Gans",
		LastName:   "Pangestu",
		Age:        24,
		IsMarried:  false,
		Hobbies:    []string{"Coding", "Gaming", "Fishing"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Bagas",
		Addresses: []Address{
			{
				Street:     "Jalan Jalan",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
			{
				Street:     "Jalan Medan Merdeka",
				Country:    "Indonesia",
				PostalCode: "12345",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street": "Jalan-Jalan", "PostalCode": "12345", "Country": "Indonesia"}, {"Street": "Jalan-Jalan2", "PostalCode": "12345", "Country": "Indonesia"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}


func TestOnlyJsonArrayComplexEncode(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Jalan",
			Country:    "Indonesia",
			PostalCode: "12345",
		},
		{
			Street:     "Jalan Medan Merdeka",
			Country:    "Indonesia",
			PostalCode: "12345",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}