package fields

import (
	"context"
	"github.com/graphql-go/graphql"
	"api/data"
	"api/types"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"fmt"
)

// GetAllProjects is a field to get a list of all projects
var GetAllProjects = &graphql.Field {
	Type:        graphql.NewList(types.ProjectType),
	Description: "Get Project List",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		all := getProjects()
		return all, nil
	},
}



func getProjects() []*types.Project {
	// get from database

	var results []*types.Project
	projectCollection := mongo.Client.Database("seedspace").Collection("projects")
	cur, err := projectCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem types.Project
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
	
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	return results
}