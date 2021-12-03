package apiv2

import "net/http"

func (a *API) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Tfp-Api-Version", "2.5")
	w.WriteHeader(http.StatusNoContent)
}
