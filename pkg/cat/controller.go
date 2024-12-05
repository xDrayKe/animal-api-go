package cat

import (
	"encoding/json"
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"

	"github.com/go-chi/render"
)

var cfg *config.Config

func InitController(config *config.Config) {
	cfg = config
}

func GetAllCats(w http.ResponseWriter, r *http.Request) {
	repo := dbmodel.NewCatRepository(cfg.DB)
	cats, err := repo.FindAll()
	if err != nil {
		http.Error(w, "Error fetching cats", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, cats)
}

func CreateCat(w http.ResponseWriter, r *http.Request) {
	var cat dbmodel.Cat
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	repo := dbmodel.NewCatRepository(cfg.DB)
	if err := repo.Create(&cat); err != nil {
		http.Error(w, "Error saving cat", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, cat)
}
