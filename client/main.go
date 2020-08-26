package main

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	oidc "github.com/coreos/go-oidc"
)

var (
	clientID     = "app-go"
	clientSecret = "8c488e1a-9773-4de5-afd1-d75b53b6e82f"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "magica"

	// redireciona para a tela de login do keycloack
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	// url de callback definida em oauth2.Config
	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not macth", http.StatusBadRequest)
			return
		}

		// authorization token
		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "faild to exchange token", http.StatusBadRequest)
			return
		}

		// authentication token
		rawIDToken, ok :=oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "no id_token", http.StatusBadRequest)
			return
		}

		resp := struct {
				OAuth2Token *oauth2.Token
				RawIDToken string
			}{
			oauth2Token, rawIDToken,
			}

		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(data)

	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}