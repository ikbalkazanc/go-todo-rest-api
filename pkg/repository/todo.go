package repository

import (
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/models"
	"github.com/jinzhu/gorm"
)

// Repository ...
type Repository struct {
	db *gorm.DB
}

// NewRepository ...
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// All ...
func (p *Repository) All() ([]models.ToDo, error) {
	todos := []models.ToDo{}
	err := p.db.Find(&todos).Error
	return todos, err
}

// FindByID ...
func (p *Repository) FindByID(id uint) (*models.ToDo, error) {
	todos := new(models.ToDo)
	err := p.db.Where(`id = ?`, id).First(&todos).Error
	return todos, err
}

// Save ...
func (p *Repository) Save(todo *models.ToDo) (*models.ToDo, error) {
	err := p.db.Save(&todo).Error
	return todo, err
}

// Delete ...
func (p *Repository) Delete(id uint) error {
	err := p.db.Delete(&models.ToDo{ID: id}).Error
	return err
}

// Migrate ...
func (p *Repository) Migrate() error {
	return p.db.AutoMigrate(&models.ToDo{}).Error
}
