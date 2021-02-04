package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	req, err := http.NewRequest("Get", "https://hive.shopbase.com/admin", nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	// Set cookie
	req.Header.Set("Cookie", fmt.Sprintf("PHPSESSID=%s", request.User))
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if res.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(res.Body)

		// error handle
		if err != nil {
			w.WriteHeader(res.StatusCode)
			_, _ = w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(res.StatusCode)
			_, _ = w.Write(data)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("tok"))
}

func main() {

	http.HandleFunc("/tick", hello)

	_ = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
