package user

import (
	"context"

	"github.com/janderfrancisco/go_bid/internal/validator"
)

type CreateUserReq struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Bio          string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval = make(validator.Evaluator)

	eval.CheckField(validator.NotBlank(req.UserName), "user_name", "This field cannot be blank")
	eval.CheckField(validator.NotBlank(req.Email), "email", "This field cannot be blank")
	eval.CheckField(
		validator.MinChars(req.Bio, 10) && validator.MaxChars(req.Bio, 200),
		"bio",
		"Bio must be between 10 and 200 characters",
	)
	eval.CheckField(validator.MinChars(req.Password, 8), "password", "Password must be at least 8 characters")
	eval.CheckField(validator.EmailRX.MatchString(req.Email), "email", "Invalid email format")

	return eval
}