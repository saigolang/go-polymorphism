package api

import "github_practice/practice/go-polymorphism-concurrency-waitgroups/structs"

// Getter will call the source API based on the user's request
// if a new source system comes in then we dont need to disturb this logic
func Getter(api []SourceAPI, request string) structs.Response {

	for _, r := range api {
		if r.source() == request {
			response := r.get()
			return response
		}
	}
	return structs.Response{}
}

type SourceAPI interface {
	get() structs.Response
	source() string
}
