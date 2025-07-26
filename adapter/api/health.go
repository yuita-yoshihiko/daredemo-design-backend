package api

import "net/http"

type HealthApi struct{}

func NewHealthApi() *HealthApi {
	return &HealthApi{}
}

func (a *HealthApi) FetchHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
