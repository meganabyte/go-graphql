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
		return projectList, nil
	},
}

var projectList = create()


func create() []types.Project {
    author := &types.Guest{Email: "megbobba@gmail.com", Name: "Megana Bobba", Projects: []int{1}}
    project := types.Project{
		ID: 1,
        Stars : 0,
		Author : *author,
		DatePosted : "August 2050",
		Title : "Marketing for small candle-making business",
		Description : "Hello world hello world hello world hello world hello world",
		Funding : 700,
		AreaOfStudy : "Business",
		IsRemote : false,
		Location: 94024,
		IsActive :true,
    }

    var projects []types.Project
    projects = append(projects, project)

    return projects
}