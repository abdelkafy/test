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
	url := "http://10.1.4.8:8080/register"
	client := &http.Client{}
	for i := 0; i < 1000000000000; i++ {
		username, email := generate()
		formData := "username-section=" + username + "&email-section=" + email + "&password-section=" + "abdo1234@" + "&password-checker-section=" + "abdo1234@" + "&register-button=submited"
		req, _ := http.NewRequest("POST", url, bytes.NewBufferString(formData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Generate a random IPv4 and set X-Forwarded-For
		req.Header.Set("X-Forwarded-For", randomIPv4())

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		resp.Body.Close()
		fmt.Println(resp)

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

func randomIPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", 1+rand.Intn(254), 1+rand.Intn(254), 1+rand.Intn(254), 1+rand.Intn(254))
}
