package service

import (
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/models"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/repository"
)

type ToDoService struct {
	ToDoRepository *repository.Repository
}

func NewToDoService(p *repository.Repository) ToDoService {
	return ToDoService{ToDoRepository: p}
}

func (p *ToDoService) All() ([]models.ToDo, error) {
	return p.ToDoRepository.All()
}

func (p *ToDoService) FindByID(id uint) (*models.ToDo, error) {

	return p.ToDoRepository.FindByID(id)
}

func (p *ToDoService) Save(todo *models.ToDo) (*models.ToDo, error) {
	return p.ToDoRepository.Save(todo)
}

func (p *ToDoService) Delete(id uint) error {
	return p.ToDoRepository.Delete(id)
}

func (p *ToDoService) Migrate() error {
	return p.ToDoRepository.Migrate()
}
