package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "http://10.1.4.4:8080/register"
	for i := 0; i <1000000000000; i++ {
		username, email := generate()
		formData := []byte("username-section=" + username + "&email-section=" + email + "&password-section=" + "abdo1234@" + "&password-checker-section=" + "abdo1234@"+"&register-button=submited")
		resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(formData))
		fmt.Println(username,"*",email)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	resp.Body.Close()
		time.Sleep(200*time.Millisecond)
	}
}


func generate() (string, string) {
	usernameLength := rand.Intn(8) + 6 
	emailLength := rand.Intn(8) + 6
	username := RandomString(usernameLength)
	email := RandomString(emailLength) + "@gmail.com"
	return username, email
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	src = rand.NewSource(time.Now().UnixNano())
	r   = rand.New(src)
	mu  sync.Mutex // prevents concurrent map writes
)

func RandomString(n int) string {
	b := make([]byte, n)
	mu.Lock()
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	mu.Unlock()
	return string(b)
}
