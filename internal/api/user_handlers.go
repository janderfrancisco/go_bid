package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/janderfrancisco/go_bid/internal/jsonutils"
	"github.com/janderfrancisco/go_bid/internal/services"
	"github.com/janderfrancisco/go_bid/internal/use_case/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[*user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(), 
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	fmt.Printf("Received valid signup data: %+v\n", data)

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]string{"user_id": id.String()})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[*user.LoginUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.AuthenticateUser(r.Context(),
		data.Email,
		data.Password,
	)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": "invalid email or password",
				"details": err.Error(),
			})
			return
		}
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected error occurred",
			"details": err.Error(),
		})
		return
	}

	// renovar o token do usuário
	errToken := api.Sessions.RenewToken(r.Context())
	if errToken != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "failed to renew user token",
			"details": errToken.Error(),
		})
		return
	}
 
	// seta o token renovado na sessão do usuário.
	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id)

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "user logged in successfully",
		"user_id": id.String(),
	})
}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "failed to log out user",
			"details": err.Error(),
		})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "user logged out successfully",
	})
}