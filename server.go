package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	rl "rate-limiter/leaky_bucket"
	u "rate-limiter/util"
)

var ch chan *u.PostRequest

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while reading request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	var pr u.PostRequest
	if err := json.Unmarshal(body, &pr); err != nil {
		http.Error(w, "Error while parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := rl.ProcessWithLimit(&pr, ch); err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusTooManyRequests)
	}
}

func main() {
	http.HandleFunc("/post", handleRequest)
	http.HandleFunc("/post/", handleRequest)

	ch = rl.Init()

	log.Println("Running server at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
