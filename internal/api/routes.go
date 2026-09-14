package api	

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (api *Api) BindRoutes() error {
	api.Router.Route("/", func(r chi.Router) {
		// Define your routes here
		r.Route("/v1", func(r chi.Router) {
			// Define your v1 routes here
			r.Route("/users/", func (r chi.Router)  {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLoginUser)
				r.Post("/logout", api.handleLogOutUser)
			})
		})


		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
	})
	return nil
} 