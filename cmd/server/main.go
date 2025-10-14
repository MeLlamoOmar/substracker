package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handleMain)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", router)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo con Air!!!!!!!"))
}