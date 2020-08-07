package fields

import (
	"github.com/graphql-go/graphql"
	//"api/data"
	"api/types"
)


// CreateProject creates a new project
var CreateProject = &graphql.Field {
	Type: types.ProjectType,
	Description: "Create a Project",
	Args: graphql.FieldConfigArgument{
	    "ID": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"Stars": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"Author": &graphql.ArgumentConfig{
			Type: types.GuestType, 
		},
		"DatePosted": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Funding": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"AreaOfStudy": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"IsRemote": &graphql.ArgumentConfig{						
			Type: graphql.Boolean,
		},
		"Location": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"IsActive": &graphql.ArgumentConfig{						
			Type: graphql.Boolean,
		},
	},
	
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		project := types.Project{
			Title: params.Args["Title"].(string),
			
		}
		//insert into db
		return project, nil
	},
}


func save() {

}