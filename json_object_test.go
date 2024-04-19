package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	IsMarried  bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Bagas",
		MiddleName: "Ganteng",
		LastName:   "Pangestu",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}
