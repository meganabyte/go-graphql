package queries

import (
	"github.com/graphql-go/graphql"
	"api/queries/fields"
)

// RootQuery is ...
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
	  "project": fields.Project,
	  "allprojects": fields.AllProjects,
	},
})


