package dbmodel

import "gorm.io/gorm"

type Cat struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Breed  string `json:"breed"`
	Weight float64 `json:"weight"`
}

func MigrateCat(db *gorm.DB) {
	db.AutoMigrate(&Cat{})
}

type CatRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) *CatRepository {
	return &CatRepository{db: db}
}

func (r *CatRepository) Create(cat *Cat) error {
	return r.db.Create(cat).Error
}

func (r *CatRepository) FindAll() ([]Cat, error) {
	var cats []Cat
	err := r.db.Find(&cats).Error
	return cats, err
}

// Implement Update and Delete similarly.
