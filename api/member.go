package api

import "github_practice/practice/go-polymorphism-concurrency-waitgroups/structs"

type Member struct {
	Name string
}

// The logic to call member API stays in this function
func (m Member) get() structs.Response {
	return structs.Response{
		FirstName: "MemberAPI",
		LastName:  "Doe",
	}
}

func (m Member) source() string {
	return m.Name
}
