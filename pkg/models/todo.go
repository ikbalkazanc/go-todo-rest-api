package models

import "time"

type ToDo struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	IsDeleted   bool      `json:"isDeleted"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"isCompleted"`
}

// PostDTO is our data transfer object for Post
type ToDoDTO struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

// ToPost converts postDTO to post
func ToToDoModel(todoDTO *ToDoDTO) *ToDo {
	return &ToDo{
		Title:       todoDTO.Title,
		IsCompleted: todoDTO.IsCompleted,
	}
}

// ToPostDTO converts post to postDTO
func ToToDoDTO(todo *ToDo) *ToDoDTO {
	return &ToDoDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		IsCompleted: todo.IsCompleted,
	}
}
