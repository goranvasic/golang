package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	rawProperties, err := os.ReadFile("assets/properties.json")
	if err != nil {
		log.Fatal(err)
	}
	properties := string(rawProperties)
	fmt.Printf("Properties: %v\n", properties)
	message := "{\n\"message\": \"Hello World\"\n}"
	rawMessage := []byte(message)
	err = os.WriteFile("assets/properties-new.json", rawMessage, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
