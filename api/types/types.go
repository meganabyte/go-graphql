package types


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

// Project is ....
type Project struct {
	Stars int
	Author Guest
	DatePosted string
	Description string
	Funding int
	AreaOfStudy string
	IsRemote bool
	Location int // area code
	IsActive bool
}

// Student is ....
type Student struct {
	Email string 
	Name string 
	Projects []Project	// starred projects
}

// Guest is ....
type Guest struct {
	Email string 
	Name string 
	Projects []Project	// opened projects
}