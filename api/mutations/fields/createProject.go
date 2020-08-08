package fields

import (
	"context"
	"github.com/graphql-go/graphql"
	"api/data"
	"api/types"
	"log"
	"fmt"
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
			// ....
		}
		createProject(project)
		return project, nil
	},
}


func createProject(project types.Project) {
	projectCollection := mongo.Client.Database("seedspace").Collection("projects")
	insertResult, err := projectCollection.InsertOne(context.TODO(), project)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}