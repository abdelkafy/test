package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://10.1.4.8:8080/login"
	for {

		resp, _ := http.Get(url)
		fmt.Println(resp)
	}
}
