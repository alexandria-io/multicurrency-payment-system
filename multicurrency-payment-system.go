package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initialize gorilla mux
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/v1/{method:[a-z]+}", MethodHandler).Methods("GET")
	http.Handle("/", rtr)

	// start listening on port 3000
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Printf("ListenAndServe:%s\n", err.Error())
	}
}

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("success")
}
