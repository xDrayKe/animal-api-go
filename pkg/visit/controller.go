package visit

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vet-clinic-api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type VisitController struct {
	repo *dbmodel.VisitRepository
}

func NewVisitController(repo *dbmodel.VisitRepository) *VisitController {
	return &VisitController{repo: repo}
}

// CreateVisit handles the creation of a new visit
func (c *VisitController) CreateVisit(w http.ResponseWriter, r *http.Request) {
	var visit dbmodel.Visit
	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.repo.Create(&visit); err != nil {
		http.Error(w, "Failed to create visit", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, visit)
}

// GetAllVisits retrieves all visits
func (c *VisitController) GetAllVisits(w http.ResponseWriter, r *http.Request) {
	visits, err := c.repo.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve visits", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, visits)
}

// GetVisitByID retrieves a visit by ID
func (c *VisitController) GetVisitByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid visit ID", http.StatusBadRequest)
		return
	}

	visit, err := c.repo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Visit not found", http.StatusNotFound)
		return
	}

	render.JSON(w, r, visit)
}

// DeleteVisit deletes a visit by ID
func (c *VisitController) DeleteVisit(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid visit ID", http.StatusBadRequest)
		return
	}

	if err := c.repo.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete visit", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
