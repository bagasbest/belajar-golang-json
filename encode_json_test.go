package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	/// UNTUK ENCODE JSON KITA PAKE json.Marshal()
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Bagas Ganteng")
	logJson([]string{"Bagas", "Ganteng"})
	logJson(1)
	logJson(true)
	logJson(nil)
}
