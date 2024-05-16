package lib

import (
	"encoding/json"
	"log"
	"os"
)

type Data struct {
	A []float32 `json:"a"`
	B []float32 `json:"b"`
}

func ParseData(filename string) (*Data, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", filename, err)
		return nil, err
	}

	var d Data
	err = json.Unmarshal(data, &d)
	if err != nil {
		log.Fatalf("Error parsing JSON data from file %s: %v", filename, err)
		return nil, err
	}

	if len(d.A) != len(d.B) {
		log.Fatalf("Error: lengths of 'a' and 'b' arrays do not match in file %s", filename)
		return nil, err
	}

	if len(d.A) < 7 || len(d.B) < 7 {
		log.Fatalf("Error: arrays 'a' and 'b' must contain at least 7 elements in file %s", filename)
		return nil, err
	}

	for i := 0; i < len(d.A); i++ {
		if d.A[i] == 0 {
			log.Fatal("a must be greater than zero")
		}
	}
	return &d, nil
}
