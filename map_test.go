package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","price": 20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id": "P0001", "name": "Apple Mac Book Pro", "price": 2000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
