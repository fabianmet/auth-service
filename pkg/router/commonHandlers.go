package router

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/fabianmet/auth-service/pkg/types"
)

//pingHandler because im playing around
func (router *Router) pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

//whoAmIHandler returns information about who you are.
func (router *Router) whoAmIHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("idtoken")

	// assume not logged in
	if err != nil {
		http.Redirect(w, r, "/iam", http.StatusFound)
		return
	}
	fmt.Println(cookie.Value)
}

//iAmHandler returns information about who you are.
func (router *Router) iAmHandler(w http.ResponseWriter, r *http.Request) {

	Groups := []string{
		"admin",
		"User",
	}

	u := &url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   "picture_goes_here.png",
	}

	User := &types.User{
		Subject:           "hoppity",
		EmailVerified:     true,
		GiveName:          "FirstName",
		FamilyName:        "LastName",
		Picture:           u,
		PreferredUserName: "Clown",
		Email:             "foo@example.com",
		Groups:            Groups,
	}

	cookie := &http.Cookie{
		Name:     "idtoken",
		Value:    string(router.Server.Key.GenerateJWT(User)),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}
