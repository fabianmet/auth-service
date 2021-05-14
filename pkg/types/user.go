package types

import "net/url"

type User struct {
	Subject           string
	EmailVerified     bool
	GiveName          string
	FamilyName        string
	Picture           *url.URL
	PreferredUserName string
	Email             string
	Groups            []Group
}

type Group struct {
	Name string
}
