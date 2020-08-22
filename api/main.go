package main

import (
	"fmt"
	"net/http"
	"api/handlers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"api/queries"
	"api/mutations"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries.RootQuery,
	Mutation: mutations.RootMutation,
})


func main() {

	h := handler.New(&handler.Config{
		Schema: &schema,
		GraphiQL:true,
		Pretty: true,
	})

	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("../client"))))
	http.HandleFunc("/", handlers.HandleMain)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/callback", handlers.HandleCallback)
	http.HandleFunc("/dashboard", handlers.HandleDashboard)
	http.Handle("/graphql", disableCors(h))
	http.HandleFunc("/logout", handlers.HandleLogout)
	fmt.Println(http.ListenAndServe(":8080", nil))

}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.WriteHeader(http.StatusOK)
		return
		}
		h.ServeHTTP(w, r)
   })
}


