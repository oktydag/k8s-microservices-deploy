package main

import (
	"customer-services/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "A customer will be added.")

	var message = ""
	message = repository.InitializeRedis()
	fmt.Fprintln(w, message)

	repository.IncrementValue("customer-count")

	var result,_ = repository.GetValue("customer-count")
	fmt.Fprintln(w, result)

	fmt.Fprintln(w, "New customer added !")

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/api/customer/add", AddCustomer)
	log.Fatal(http.ListenAndServe(":5005", router))
}