package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/models"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/service"
)

type ToDoAPI struct {
	ToDoService service.ToDoService
}

func NewToDoAPI(p service.ToDoService) ToDoAPI {
	return ToDoAPI{ToDoService: p}
}

// FindAll
func (p ToDoAPI) FindAllTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		todos, err := p.ToDoService.All()
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, todos)
	}
}

// FindByID ...
func (p ToDoAPI) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if id is integer
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Find login by id from db
		todo, err := p.ToDoService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, models.ToToDoDTO(todo))
	}
}

// CreatePost ...
func (p ToDoAPI) CreateToDo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var todoDto models.ToDoDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&todoDto); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdtodo, err := p.ToDoService.Save(models.ToToDoModel(&todoDto))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, models.ToToDoDTO(createdtodo))
	}
}

func (p ToDoAPI) DeleteToDo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		todo, err := p.ToDoService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = p.ToDoService.Delete(todo.ID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		response := Response{
			Message: "ToDo deleted successfully!",
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}

// Migrate ...
func (p ToDoAPI) Migrate() {
	err := p.ToDoService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
