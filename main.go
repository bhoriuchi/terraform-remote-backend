package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/.well-known/terraform.json", handleTFJSON)
	http.ListenAndServeTLS(":9443", "dev-cert.pem", "dev-key.pem", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/.well-known/terraform.json" {
			handleTFJSON(w, r)
		} else {
			// fmt.Println(r)
			fmt.Println(r.Method, "", r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
}

func handleTFJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{
		"modules.v1":"/api/registry/v1/modules/",
		"motd.v1":"/api/terraform/motd",
		"state.v2":"/api/v2/",
		"tfe.v2":"/api/v2/",
		"tfe.v2.1":"/api/v2/",
		"tfe.v2.2":"/api/v2/",
		"versions.v1":"https://checkpoint-api.hashicorp.com/v1/versions/",
		"login.v1": {
			"client": "terraform-server",
			"authz": "http://localhost:5556/auth",
			"token": "http://localhost:5556/token",
			"scopes": [
				"groups", "openid", "email", "profile", "offline_access"
			]
		}
	}`))
}
