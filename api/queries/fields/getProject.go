package fields

import (
	"github.com/graphql-go/graphql"
	//"go.mongodb.org/mongo-driver/bson"	
	//"api/data"
	"api/types"
)

// Project is a field to get a project by ID
var Project = &graphql.Field {
	Type:        types.ProjectType,
	Description: "Get Project By ID",
	Args: graphql.FieldConfigArgument{
		"ID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// get project from database
		return nil, nil
	},
}

func add() {

}