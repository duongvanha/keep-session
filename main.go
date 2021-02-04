package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	User string `json:"user"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	client := &http.Client{}

	// Declare HTTP Method and Url
	req, err := http.NewRequest("Get", "https://en7nthov5hqb.x.pipedream.net/", nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	// Set cookie
	req.Header.Set("Cookie", fmt.Sprintf("PHPSESSID=%s", request.User))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err = client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("tok"))
}

func main() {

	http.HandleFunc("/tick", hello)

	_ = http.ListenAndServe(":8080", nil)
}
