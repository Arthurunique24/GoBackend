package main

import (
	"encoding/json"
	"fmt"
	"log"
	"awesomeProject/JSON/EasyJSON/test"
)

func main() {
	var data test.Schema
	if err := readJson(&data); err != nil {
		log.Fatalf("can't read json data: %s", err)
	}

	dataToPrint, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("%s", dataToPrint)
}
