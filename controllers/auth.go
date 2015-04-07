package controllers

import (
	//"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	//"jaabo/models"
	"net/http"
)

func GoogleAuthHandler(w http.ResponseWriter, r *http.Request) {
	conf := &oauth2.Config{
		ClientID:     "1021671356483-kp0rapcbaqua3vq26u5d0pph2jpdel2o.apps.googleusercontent.com",
		ClientSecret: "lSPKn6BryeF7C6gOEu9eloOx",
		Scopes: []string{
			"openid",
			"email",
		},
		RedirectURL: "http://localhost:8080/auth/handle/google",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	authcode := r.FormValue("code")

	tok, err := conf.Exchange(oauth2.NoContext, authcode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", tok)
	//fmt.Printf("%+v\n", (*tok).id_token)

	//create session variable
	//set the cookie

}
