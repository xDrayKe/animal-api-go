package treatment

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vet-clinic-api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TreatmentController struct {
	repo *dbmodel.TreatmentRepository
}

func NewTreatmentController(repo *dbmodel.TreatmentRepository) *TreatmentController {
	return &TreatmentController{repo: repo}
}

// CreateTreatment handles the creation of a new treatment
func (c *TreatmentController) CreateTreatment(w http.ResponseWriter, r *http.Request) {
	var treatment dbmodel.Treatment
	if err := json.NewDecoder(r.Body).Decode(&treatment); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.repo.Create(&treatment); err != nil {
		http.Error(w, "Failed to create treatment", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, treatment)
}

// GetAllTreatments retrieves all treatments
func (c *TreatmentController) GetAllTreatments(w http.ResponseWriter, r *http.Request) {
	treatments, err := c.repo.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve treatments", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, treatments)
}

// GetTreatmentByID retrieves a treatment by ID
func (c *TreatmentController) GetTreatmentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid treatment ID", http.StatusBadRequest)
		return
	}

	treatment, err := c.repo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Treatment not found", http.StatusNotFound)
		return
	}

	render.JSON(w, r, treatment)
}

// DeleteTreatment deletes a treatment by ID
func (c *TreatmentController) DeleteTreatment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid treatment ID", http.StatusBadRequest)
		return
	}

	if err := c.repo.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete treatment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
