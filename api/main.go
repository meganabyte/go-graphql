package main

import (
	"fmt"
	"net/http"
	"api/oauth"
)

// NotImplemented message will be returned if handler not done
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Not Implemented"))
})  


func main() {
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("../client"))))
	http.HandleFunc("/", oauth.HandleMain)
	http.HandleFunc("/login", oauth.HandleLogin)
	http.HandleFunc("/callback", oauth.HandleCallback)
	http.HandleFunc("/dashboard", NotImplemented)
	http.HandleFunc("/dashboard/starred", NotImplemented)
	http.HandleFunc("/logout", NotImplemented)

	fmt.Println(http.ListenAndServe(":8080", nil))
}

