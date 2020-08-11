package fields

import (
	"github.com/graphql-go/graphql"
	"api/data"
	"api/types"
	"context"
	"log"
	"fmt"
)

// CreateStudent creates a new project
var CreateStudent = &graphql.Field {
	Type: types.StudentType,
	Description: "Create a Student user",
	Args: graphql.FieldConfigArgument{
	    "Email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Projects": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int), 
		},
	},
	
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		student := types.Student{
			Email: params.Args["Email"].(string),
			// ....
		}
		createStudent(student)
		return student, nil
	},
}


func createStudent(student types.Student) {
	studentCollection := mongo.Client.Database("seedspace").Collection("students")
	insertResult, err := studentCollection.InsertOne(context.TODO(), student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}