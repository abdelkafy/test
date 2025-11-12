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
	client := &http.Client{}
	url := "http://10.1.4.4:8080/like-post"

	for  {
		title, _ := generatePost()
		
		formData := []byte("post_id=" +"1" + "&value=" + "1")

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(formData))
		if err != nil {
			panic(err)
		}

		// Set headers
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// ✅ Add your session cookie here
		req.AddCookie(&http.Cookie{
			Name:  "id",       // change this to match your actual cookie name
			Value: "d628f958-dd2d-4cb6-8d3a-95ca53b4607a", // replace this with your valid session ID
		})

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Created post:", title)
		resp.Body.Close()
		time.Sleep(200 * time.Millisecond)
	}
}

func generatePost() (string, string) {
	titleLength := rand.Intn(8) + 6
	contentLength := rand.Intn(20) + 20
	title := RandomString(titleLength)
	content := RandomString(contentLength)
	return title, content
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	src = rand.NewSource(time.Now().UnixNano())
	r   = rand.New(src)
	mu  sync.Mutex
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
