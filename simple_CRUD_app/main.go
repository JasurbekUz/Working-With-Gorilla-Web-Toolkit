package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

const PORT = ":4000"

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/classifieds", GetClassifiedsCtrl).Methods("GET")
	router.HandleFunc("/classifieds/{id}", GetClassifiedCtrl).Methods("GET")
	//router.HandleFunc("/classified/search", SearchGetClassifiedCtrl).Methods("GET").Queries("q", "{id}")
	router.HandleFunc("/classified/post", PostClassifiedsCtrl).Methods("POST")
	router.HandleFunc("/classified/put", PutClassifiedsCtrl).Methods("PUT")
	router.HandleFunc("/classified/rm", DelClassifiedCtrl).Methods("DELETE")

	log.Printf("server is ready at %s ", PORT)

	err := http.ListenAndServe(PORT, router)

	log.Fatalf("server error %v", err)
}
