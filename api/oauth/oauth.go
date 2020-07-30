package oauth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
	"html/template"
	"encoding/json"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString = "pseudo-random"
)

// User is a retrieved & authenticated user
type User struct {
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


func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// HandleMain serves the basic HTML for the landing page
func HandleMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../client/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// HandleLogin gets a url based on pseudo-random state that requests the defined scopes
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleCallback returns the requested user info and prints it
func HandleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	u := User{}
	if err = json.Unmarshal(content, &u); err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// lookup user in database
		// if not in database, create new user from google data
	// return user and log them in 
	fmt.Println(u)
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("oauth state is not valid")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("could not get access token: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("could not create get request: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}