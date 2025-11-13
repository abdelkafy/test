package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	url := "http://10.1.4.8:8080/login"
	client := &http.Client{}
	for {
		req,err:=http.NewRequest("GET",url,nil)
		if err!=nil {
			fmt.Println(err)
		}
		req.Header.Set("X-Forwarded-For",randomIPv4())
		resp, err := client.Do(req)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	}
}

func randomIPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", 1+rand.Intn(254), 1+rand.Intn(254), 1+rand.Intn(254), 1+rand.Intn(254))
}
