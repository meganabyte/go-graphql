package fields

import (
	//"context"
	"github.com/graphql-go/graphql"
	//"github.com/mongodb/mongo-go-driver/bson"	
	//"api/data"
	"api/types"
)

// AllProjects is a field to get a list of all projects
var AllProjects = &graphql.Field {
	Type:        graphql.NewList(types.ProjectType),
	Description: "Get Project List",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// get from database
		return nil, nil
	},
}
