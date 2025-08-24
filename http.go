package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func json_example() {
	raw := []byte(`{"price": 10.99, "title": "Book"}`)

	var result map[string]interface{}
	json.Unmarshal(raw, &result)

	// Extracting float64
	price, ok := getValue[float64](result, "price")
	if !ok {
		fmt.Println("Could not get price")
		return
	}

	// Extracting string
	title, ok := getValue[string](result, "title")
	if !ok {
		fmt.Println("Could not get title")
		return
	}

	fmt.Printf("Title: %s - Price: %.2f\n", title, price)

}

// httpGet performs an HTTP GET request and returns the response body as a string
func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
