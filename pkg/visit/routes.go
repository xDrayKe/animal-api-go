package visit

import (
	"vet-clinic-api/pkg/authentication"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, controller *VisitController) {
	r.Route("/visits", func(r chi.Router) {
		r.Use(authentication.AuthMiddleware) // Ajout du middleware d'authentification
		r.Post("/", controller.CreateVisit)   // Créer une consultation
		r.Get("/", controller.GetAllVisits)   // Obtenir toutes les consultations
		r.Get("/{id}", controller.GetVisitByID) // Obtenir une consultation par ID
		r.Put("/{id}", controller.UpdateVisit) // Mettre à jour une consultation
		r.Delete("/{id}", controller.DeleteVisit) // Supprimer une consultation
	})
}
