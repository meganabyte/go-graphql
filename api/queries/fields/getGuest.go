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

// Guest is a field to get a project by email
var Guest = &graphql.Field {
	Type:        types.GuestType,
	Description: "Get Guest By email",
	Args: graphql.FieldConfigArgument{
		"Email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var result types.Guest
		email, ok := p.Args["Email"].(string)
		if ok {
			filter := bson.D{{Key:"email", Value:email}}
			studentCollection := mongo.Client.Database("seedspace").Collection("guests")
			err := studentCollection.FindOne(context.TODO(), filter).Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Found a single document: %+v\n", result)
			return result, nil
		}
		return nil, nil
	},
}