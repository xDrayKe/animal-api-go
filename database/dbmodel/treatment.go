package dbmodel

import "gorm.io/gorm"

type Treatment struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"` // Nom du traitement ou médicament
	Dosage     string `json:"dosage"` // Dosage prescrit
	Instructions string `json:"instructions"` // Instructions pour le traitement
	VisitID    uint   `json:"visit_id"` // ID de la consultation associée
	Visit      Visit  `gorm:"foreignKey:VisitID"` // Relation avec le modèle Visit
}

func MigrateTreatment(db *gorm.DB) {
	db.AutoMigrate(&Treatment{})
}

type TreatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRepository(db *gorm.DB) *TreatmentRepository {
	return &TreatmentRepository{db: db}
}

func (r *TreatmentRepository) Create(treatment *Treatment) error {
	return r.db.Create(treatment).Error
}

func (r *TreatmentRepository) FindAll() ([]Treatment, error) {
	var treatments []Treatment
	err := r.db.Preload("Visit").Find(&treatments).Error
	return treatments, err
}

func (r *TreatmentRepository) FindByID(id uint) (*Treatment, error) {
	var treatment Treatment
	err := r.db.Preload("Visit").First(&treatment, id).Error
	return &treatment, err
}

func (r *TreatmentRepository) Delete(id uint) error {
	return r.db.Delete(&Treatment{}, id).Error
}
