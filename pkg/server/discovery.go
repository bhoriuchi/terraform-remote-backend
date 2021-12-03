package server

import "net/http"

func (s *Server) handleDiscovery(w http.ResponseWriter, r *http.Request) {
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
