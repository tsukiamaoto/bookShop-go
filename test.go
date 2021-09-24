package main

import (
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`my first website`))
}

func main() {
	// router setting
	http.HandleFunc("/", test)

	// start server
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
