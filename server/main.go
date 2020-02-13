package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("go ran")

	http.HandleFunc("index.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
	})

	http.HandleFunc("components.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
	})

	http.HandleFunc("*.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("everything else")
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server Started")
	}

}
