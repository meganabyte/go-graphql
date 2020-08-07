package mutations

import (
	"github.com/graphql-go/graphql"
	"api/mutations/fields"
)

// RootMutation is the entry point for graphql mutations
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createProject": fields.CreateProject,
		"createGuest": fields.CreateGuest,
		"createStudent": fields.CreateStudent,
	},
})