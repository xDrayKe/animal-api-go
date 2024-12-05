package cat

import (
	"vet-clinic-api/pkg/authentication"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, controller *CatController) {
	r.Route("/cats", func(r chi.Router) {
		r.Use(authentication.AuthMiddleware) // Ajout du middleware d'authentification
		r.Post("/", controller.CreateCat)    // Créer un chat
		r.Get("/", controller.GetAllCats)    // Obtenir tous les chats
		r.Get("/{id}", controller.GetCatByID) // Obtenir un chat par ID
		r.Put("/{id}", controller.UpdateCat) // Mettre à jour un chat
		r.Delete("/{id}", controller.DeleteCat) // Supprimer un chat
	})
}
