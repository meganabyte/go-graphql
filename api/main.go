package main

import (
	"fmt"
	//"net/http"
	//"api/handlers"
	"github.com/graphql-go/graphql"
	"log"
	"encoding/json"
	"api/queries"
	"api/mutations"
)

/*

// NotImplemented message will be returned if handler not done
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Not Implemented"))
})  

*/

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries.RootQuery,
	Mutation: mutations.RootMutation,
})


func main() {
	/*

	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("../client"))))
	http.HandleFunc("/", handlers.HandleMain)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/callback", handlers.HandleCallback)
	http.HandleFunc("/dashboard", handlers.HandleDashboard)
	http.HandleFunc("/dashboard/starred", NotImplemented)
	http.HandleFunc("/logout", handlers.HandleLogout)
	fmt.Println(http.ListenAndServe(":8080", nil))

	*/


	// Query or mutation
	request := `
			mutation {
				createProject(Title: "Hello World") {
					Title
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

}


