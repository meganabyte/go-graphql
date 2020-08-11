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
		var result types.Project
		id, ok := p.Args["ID"].(int)
		if ok {
			filter := bson.D{{Key:"id", Value:id}}
			projectCollection := mongo.Client.Database("seedspace").Collection("projects")
			err := projectCollection.FindOne(context.Background(), filter).Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Found a single document: %+v\n", result)
			return result, nil
		}
		return nil, nil
	},
}

