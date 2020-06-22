package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url    string = "http://127.0.0.1:8080/videos"
	method string = "GET"
)

func main() {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic bG9jYWx3b3JrZXI6bG9jYWx3b3JrZXIxMjM=")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
