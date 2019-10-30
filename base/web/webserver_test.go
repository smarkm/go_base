package web_test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestWebServer(t *testing.T) {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header { // for headers
			log.Println(k, ":", v)
		}
		fmt.Fprintf(w, "hell from server")
	})

	log.Println("server started")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Printf("err:", err)
		return
	}
}
