package apiv2

import "net/http"

func (a *API) HandleOrganizationSet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}