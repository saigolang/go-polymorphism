package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	api2 "github_practice/practice/go-polymorphism-concurrency-waitgroups/api"
	"github_practice/practice/go-polymorphism-concurrency-waitgroups/structs"
	"net/http"
)

func main() {

	routes := mux.NewRouter()

	routes.HandleFunc("/getdata", handler).Methods(http.MethodPost)

	fmt.Println("started server on 8080")

	http.ListenAndServe(":8080", routes)

}

func handler(rw http.ResponseWriter, rq *http.Request) {

	var request structs.Request
	json.NewDecoder(rq.Body).Decode(&request)

	provider := api2.Provider{Name: "providerAPI"}
	member := api2.Member{Name: "memberAPI"}

	// passing our source system types to interface
	// if the new source system comes in then we will add that to interface
	api := []api2.SourceAPI{provider, member}

	// calling our getter function to decide which sourcessystem to call
	// based on the given request
	response := api2.Getter(api, request.Name)

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(&response)
}
