package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	url    string = "http://127.0.0.1:8080/videos"
	method string = "POST"
)

func main() {

	payload := strings.NewReader("{\n        \"title\": \"Star Wars: Phantom menace\",\n        \"desc\": \"Sci-Fi Adventure\",\n        \"url\": \"www.some_url_sw_last_jedi.com\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic bG9jYWx3b3JrZXI6bG9jYWx3b3JrZXIxMjM=")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
