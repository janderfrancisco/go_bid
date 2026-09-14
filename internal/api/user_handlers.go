package api

import (
	"net/http"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup User"))
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login User"))
}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Log Out User"))
}