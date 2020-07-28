package main

import (
	"fmt"
	"net/http"
	"api/oauth"
)

func main() {
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("../client"))))
	http.HandleFunc("/", oauth.HandleMain)
	http.HandleFunc("/login", oauth.HandleLogin)
	http.HandleFunc("/callback", oauth.HandleCallback)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

