package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	result := getHttpRequest("https://example.com")
	fmt.Println(result)
}

func getHttpRequest(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		rawBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		return string(rawBody)
	}
	return ""
}
