package main

import (
	"fmt"
	"net/http"
	"api/handlers"
	/*
	"github.com/graphql-go/graphql"
	"log"
	"encoding/json"
	"api/queries"
	"api/mutations"
	*/
)

/*
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries.RootQuery,
	Mutation: mutations.RootMutation,
})
*/


func main() {

	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("../client"))))
	http.HandleFunc("/", handlers.HandleMain)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/callback", handlers.HandleCallback)
	http.HandleFunc("/dashboard", handlers.HandleDashboard)
	http.HandleFunc("/logout", handlers.HandleLogout)
	fmt.Println(http.ListenAndServe(":8080", nil))

	/*
	// Query or mutation
	request := `
			{
				student(Email:"test@gmail.com") {
					Email
				}
			}
	`
	params := graphql.Params{Schema: schema, RequestString: request}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
	*/

}


