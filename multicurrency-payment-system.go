package mucupa

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test success")
}

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("quote success")
}

func MuxInit() *mux.Router {

	// create a new router with the specified listener functions
	rtr := mux.NewRouter()
	rtr.HandleFunc("/test", MethodHandler).Methods("GET")
	rtr.HandleFunc("/quote", QuoteHandler).Methods("GET")
	http.Handle("/", rtr)

	// start listening on port 3000
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Printf("ListenAndServe:%s\n", err.Error())
	}
	return rtr
}
