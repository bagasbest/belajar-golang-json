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

func TestJsonTags(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	/// MISAL KITA GAK MAU Id, Name, ImageURL besar diawal katanya maka tambahin tag jadi snake_case
	fmt.Println(string(bytes))
}

func TestJsonDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","Image_URL":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
