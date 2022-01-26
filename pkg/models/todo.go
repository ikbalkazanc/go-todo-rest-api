package models

import "time"

type ToDo struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	IsDeleted   bool      `json:"isDeleted"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"isCompleted"`
}

// TodoDTO is our data transfer object for Post
type ToDoDTO struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

// To ToDo converts todoDTO to post
func ToToDoModel(todoDTO *ToDoDTO) *ToDo {
	return &ToDo{
		Title:       todoDTO.Title,
		IsCompleted: todoDTO.IsCompleted,
	}
}

// ToTodoDTO converts post to todoDTO
func ToToDoDTO(todo *ToDo) *ToDoDTO {
	return &ToDoDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		IsCompleted: todo.IsCompleted,
	}
}
