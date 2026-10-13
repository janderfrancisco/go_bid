package api

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
)

func (api *Api) BindRoutes() error {

	api.Router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, api.Sessions.LoadAndSave)

	api.Router.Route("/", func(r chi.Router) {
		// Define your routes here
		r.Route("/api/v1", func(r chi.Router) {
			// Define your v1 routes here
			r.Route("/users/", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLoginUser)
				r.With(api.AuthMiddleware).Post("/logout", api.handleLogOutUser)
			})
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
	})
	return nil
}
