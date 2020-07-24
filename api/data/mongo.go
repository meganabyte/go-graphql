package mongo

import (
	//"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    //"go.mongodb.org/mongo-driver/mongo/readpref"
)

//var Client, err = mongo.Connect(context.Background(), "Insert your MongoDB URI here!", nil)

var Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
