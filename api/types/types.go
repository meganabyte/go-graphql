package types

import (
	"github.com/graphql-go/graphql"
)

// GoogleUser is a retrieved & authenticated Google user
type GoogleUser struct {
    ID string `json:"id"`
    Email string `json:"email"`
    VerifiedEmail bool `json:"verified_email"`
    Name string `json:"name"`
    FirstName string `json:"given_name"`
    LastName string `json:"family_name"`
    Picture string `json:"picture"`
    Locale string `json:"locale"`
    HD string `json:"hd"`
}

// Project is a seedspace project
type Project struct {
	ID int
	Stars int
	Author Guest
	DatePosted string
	Title string
	Description string
	Funding int
	AreaOfStudy string
	IsRemote bool
	Location int // area code
	IsActive bool
}


// Student is a seedspace student account
type Student struct {
	Email string 
	Name string 
	Projects []Project	// starred projects
}


// Guest is is a seedspace guest account
type Guest struct {
	Email string 
	Name string 
	Projects []int	// opened projects
}

// ProjectType is a graphql type for Project
var ProjectType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Project",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Stars": &graphql.Field{
				Type: graphql.Int,
			},
			"Author": &graphql.Field{
				Type: GuestType, 
			},
			"DatePosted": &graphql.Field{
				Type: graphql.String,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Funding": &graphql.Field{
				Type: graphql.Int,
			},
			"AreaOfStudy": &graphql.Field{
				Type: graphql.String,
			},
			"IsRemote": &graphql.Field{						
				Type: graphql.Boolean,
			},
			"Location": &graphql.Field{
				Type: graphql.Int,
			},
			"IsActive": &graphql.Field{						
				Type: graphql.Boolean,
			},
		},
	},
)

// GuestType is a graphql type for Guest
var GuestType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Guest",
		Fields: graphql.Fields{
			"Email": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Projects": &graphql.Field{			
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

// StudentType is a graphql type for Student
var StudentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"Email": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Projects": &graphql.Field{			
				Type: graphql.NewList(ProjectType),
			},
		},
	},
)