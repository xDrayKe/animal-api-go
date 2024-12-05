package treatment

import (
	"vet-clinic-api/pkg/authentication"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, controller *TreatmentController) {
	r.Route("/treatments", func(r chi.Router) {
		r.Use(authentication.AuthMiddleware)
		r.Post("/", controller.CreateTreatment)
		r.Get("/", controller.GetAllTreatments)
		r.Get("/{id}", controller.GetTreatmentByID)
		r.Delete("/{id}", controller.DeleteTreatment)
	})
}
