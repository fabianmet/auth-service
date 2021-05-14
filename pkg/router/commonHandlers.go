package router

import (
	"fmt"
	"net/http"
)

//pingHandler because im playing around
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

//whoAmIHandler returns information about who you are.
func whoAmIHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("idtoken")

	// assume not logged in
	if err != nil {
		http.Redirect(w, r, "/iam", http.StatusFound)
		return
	}
	fmt.Println(cookie.Value)
}

//iAmHandler returns information about who you are.
func iAmHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "idtoken",
		Value:    "token goes here",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}