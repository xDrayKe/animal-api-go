package dbmodel

import "gorm.io/gorm"

type Visit struct {
	ID          uint   `gorm:"primaryKey"`
	Date        string `json:"date"` // Date de la consultation
	Motif       string `json:"motif"` // Motif de la consultation
	Veterinarian string `json:"veterinarian"` // Vétérinaire en charge
	CatID       uint   `json:"cat_id"` // ID du chat concerné
	Cat         Cat    `gorm:"foreignKey:CatID"` // Relation avec le modèle Cat
}

func MigrateVisit(db *gorm.DB) {
	db.AutoMigrate(&Visit{})
}

type VisitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) *VisitRepository {
	return &VisitRepository{db: db}
}

func (r *VisitRepository) Create(visit *Visit) error {
	return r.db.Create(visit).Error
}

func (r *VisitRepository) FindAll() ([]Visit, error) {
	var visits []Visit
	err := r.db.Preload("Cat").Find(&visits).Error
	return visits, err
}

func (r *VisitRepository) FindByID(id uint) (*Visit, error) {
	var visit Visit
	err := r.db.Preload("Cat").First(&visit, id).Error
	return &visit, err
}

func (r *VisitRepository) Delete(id uint) error {
	return r.db.Delete(&Visit{}, id).Error
}
