package fields

import (
	//"context"
	"github.com/graphql-go/graphql"
	//"github.com/mongodb/mongo-go-driver/bson"	
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
		// GET PROJECT FROM DATABASE
		id, ok := p.Args["ID"].(int)
		if ok {
			// Find project
			for _, project := range allProjects {
				if int(project.ID) == id {
					return project, nil
				}
			}
		}
		return nil, nil
	},
}

var allProjects = populate()


func populate() []types.Project {
    author := &types.Guest{Email: "megbobba@gmail.com", Name: "Megana Bobba"}
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