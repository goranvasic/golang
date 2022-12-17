package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing command line arguments.")
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
