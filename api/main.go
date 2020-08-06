package main

import (
	"fmt"
	"net/http"
	"api/handlers"
	"api/types"
)

// NotImplemented message will be returned if handler not done
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Not Implemented"))
})  


func main() {
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("../client"))))
	http.HandleFunc("/", handlers.HandleMain)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/callback", handlers.HandleCallback)
	http.HandleFunc("/dashboard", handlers.HandleDashboard)
	http.HandleFunc("/dashboard/starred", NotImplemented)
	http.HandleFunc("/logout", handlers.HandleLogout)
	fmt.Println(http.ListenAndServe(":8080", nil))

	projects := populate()
	


}

func populate() []types.Project {
    author := &types.Guest{Email: "megbobba@gmail.com", Name: "Megana Bobba"}
    project := types.Project{
		ID: 1,
        Stars : 0,
		Author : *author,
		DatePosted : "August 2050",
		Title : "Marketing for small candle-making business",
		Description : "Hello world hello world hello world hello world hello world",
		Funding : 700,
		AreaOfStudy : "Business",
		IsRemote : false,
		Location: 94024,
		IsActive :true,
    }

    var projects []types.Project
    projects = append(projects, project)

    return projects
}
