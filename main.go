package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Session struct {
	Domain   string `json:"domain"`
	HostOnly bool   `json:"hostOnly"`
	HttpOnly bool   `json:"httpOnly"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	SameSite string `json:"sameSite"`
	Secure   bool   `json:"secure"`
	Session  bool   `json:"session"`
	StoreId  string `json:"storeId"`
	Value    string `json:"value"`
}

type Request struct {
	User Session `json:"user"`
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
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/admin", request.User.Domain), nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	req.Header.Set("Cookie", fmt.Sprintf("PHPSESSID=%s", request.User.Value))

	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	defer res.Body.Close()

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
