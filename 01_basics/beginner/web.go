package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", sayhello)
	http.HandleFunc("/info", myinf)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println("URL isn't correct")
		fmt.Fprintf(w, "Такого адресса нет")
	} else {
		fmt.Fprintf(w, "Hello!")
	}
}

func myinf(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Maxim Burnov")
}
