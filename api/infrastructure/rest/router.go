package rest

import (
	"github.com/go-chi/chi/v5"
	"loom-api/api/infrastructure/rest/doc"
	"loom-api/api/infrastructure/rest/sportlink/team"
	"net/http"
)

func NewRouter(tHandler team.Handler, dHandler doc.Handler) http.Handler {
	r := chi.NewRouter()
	r.Post("/sportlink/team_creation_scenario", tHandler.TeamCreationScenario)
	r.Get("/docs", dHandler.HandleDocRequest)
	r.Get("/docs/{filename}", dHandler.HandleDocRequest)

	return r
}
