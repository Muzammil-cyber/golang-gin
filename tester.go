package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// printResponse prints the HTTP status, pretty-prints JSON, or prints raw body if not JSON
func printResponse(section string, res *http.Response, body []byte, err error) {
	fmt.Printf("\n===== %s =====\n", section)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Status: %s\n", res.Status)
	if len(body) == 0 {
		fmt.Println("(Empty response body)")
		return
	}
	if json.Valid(body) {
		var jsonObj interface{}
		if err := json.Unmarshal(body, &jsonObj); err == nil {
			formatted, _ := json.MarshalIndent(jsonObj, "", "  ")
			fmt.Println(string(formatted))
			return
		}
	}
	fmt.Println(string(body))
}

func main() {

	url := "http://localhost:8080/videos"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		printResponse("GET /videos", nil, nil, err)
		return
	}
	req.Header.Add("Authorization", "Basic YWRtaW46cGFzc3dvcmQ=")

	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	body, _ := io.ReadAll(res.Body)
	printResponse("GET /videos", res, body, err)

	url = "http://localhost:8080/videos"
	newID := time.Now().UnixNano()
	payloadStr := fmt.Sprintf(`{"id": "idx_%d", "title": "First", "description": "some words", "url": "https://link.com", "author": {"name": "Cornell Sanders", "age": 30, "email": "email@gmail.co"}}`, newID)
	payload := strings.NewReader(payloadStr)
	req, err = http.NewRequest("POST", url, payload)
	if err != nil {
		printResponse("POST /videos", nil, nil, err)
		return
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "CLI with Go Script")
	req.Header.Add("Authorization", "Basic YWRtaW46cGFzc3dvcmQ=")
	req.Header.Add("Content-Type", "application/json")

	res, err = http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	body, _ = io.ReadAll(res.Body)
	printResponse("POST /videos", res, body, err)

}
