package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJsonEncode(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Bagas",
		MiddleName: "Ganteng",
		LastName:   "Pangestu",
	}

	_ = encoder.Encode(customer)
	fmt.Println(customer)
}
