package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/janderfrancisco/go_bid/internal/services"
	"github.com/alexedwards/scs/v2"
)

type Api struct {
	Router      *chi.Mux
	UserService *services.UserService
	Sessions    *scs.SessionManager
}
