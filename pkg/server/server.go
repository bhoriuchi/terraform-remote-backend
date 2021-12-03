package server

import (
	"fmt"
	"net/http"

	apiv2 "github.com/bhoriuchi/terraform-remote-backend/pkg/api/v2"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

// Server serves backend
type Server struct {
	v  *viper.Viper
	r  chi.Router
	v2 *apiv2.API
}

// NewServer creates a new server instance
func NewServer(v *viper.Viper) (*Server, error) {
	v2api, err := apiv2.NewAPI()
	if err != nil {
		return nil, err
	}

	s := &Server{
		v:  v,
		r:  chi.NewRouter(),
		v2: v2api,
	}

	return s, nil
}

// Start starts the server
func (s *Server) Start() {
	s.r.Get("/.well-known/terraform.json", s.handleDiscovery)

	s.r.Route("/api", func(r chi.Router) {
		r.Route("/v2", func(r chi.Router) {
			r.Get("/ping", s.v2.HandlePing)
			r.Route("/organizations/{org}", func(r chi.Router) {
				r.Get("/entitlement-set", s.v2.HandleOrganizationSet)
			})
		})
	})

	// catch all
	s.r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, "", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	// start listening
	fmt.Println("Listening")
	http.ListenAndServeTLS(
		s.v.GetString("http.address"),
		"dev-cert.pem",
		"dev-key.pem",
		s.r,
	)
}
