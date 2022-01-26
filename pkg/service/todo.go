package service

import (
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/models"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/repository"
)

// PostService ...
type ToDoService struct {
	ToDoRepository *repository.Repository
}

// NewPostService ...
func NewToDoService(p *repository.Repository) ToDoService {
	return ToDoService{ToDoRepository: p}
}

// All ...
func (p *ToDoService) All() ([]models.ToDo, error) {
	return p.ToDoRepository.All()
}

// FindByID ...
func (p *ToDoService) FindByID(id uint) (*models.ToDo, error) {
	return p.ToDoRepository.FindByID(id)
}

// Save ...
func (p *ToDoService) Save(post *models.ToDo) (*models.ToDo, error) {
	return p.ToDoRepository.Save(post)
}

// Delete ...
func (p *ToDoService) Delete(id uint) error {
	return p.ToDoRepository.Delete(id)
}

// Migrate ...
func (p *ToDoService) Migrate() error {
	return p.ToDoRepository.Migrate()
}
