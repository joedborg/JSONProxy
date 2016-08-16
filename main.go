package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "http://127.0.0.1:8000/api"

func requestGet(p string) []byte {
	resp, _ := http.Get(baseURL + p)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func requestPost(p string, f map[string][]string) []byte {
	resp, _ := http.PostForm(baseURL+p, f)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func receiveRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		body := requestGet(r.URL.Path)
		if len(r.FormValue("callback")) > 0 {
			fmt.Fprintf(w, "%s(%s)", r.FormValue("callback"), body)
		} else {
			w.Write(body)
		}

	case "POST":
		body := requestPost(r.URL.Path, r.Form)
		if len(r.FormValue("callback")) > 0 {
			fmt.Fprintf(w, "%s(%s)", r.FormValue("callback"), body)
		} else {
			w.Write(body)
		}
	}
}

func main() {
	http.HandleFunc("/", receiveRequest)
	http.ListenAndServe(":8080", nil)
}
