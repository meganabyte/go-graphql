package fields

import (
	"github.com/graphql-go/graphql"
	"api/data"
	"api/types"
	"context"
	"log"
	"fmt"
)

// CreateGuest creates a new project
var CreateGuest = &graphql.Field {
	Type: types.GuestType,
	Description: "Create a Guest user",
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
		guest := types.Guest{
			Email: params.Args["Email"].(string),
			// ....
		}
		createGuest(guest)
		return guest, nil
	},
}


func createGuest(guest types.Guest) {
	guestCollection := mongo.Client.Database("seedspace").Collection("guests")
	insertResult, err := guestCollection.InsertOne(context.TODO(), guest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}