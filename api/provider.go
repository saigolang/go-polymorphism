package api

import (
	"github_practice/practice/go-polymorphism-concurrency-waitgroups/structs"
)

type Provider struct {
	Name string
}

// The logic to call provider API stays in this function
func (p Provider) get() structs.Response {
	return structs.Response{
		FirstName: "ProviderAPI",
		LastName:  "Bar",
	}
}

func (p Provider) source() string {
	return p.Name
}
