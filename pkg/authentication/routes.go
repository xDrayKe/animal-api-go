package authentication

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, controller *AuthController) {
	r.Post("/login", controller.Login)
}
