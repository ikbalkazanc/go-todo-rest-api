package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/api"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/repository"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {

	a := App{}

	a.initialize()

	a.routes()

	a.run(":3001")
}

func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initialize() {
	var err error

	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "xl3236825", "kartelam")

	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) routes() {
	postAPI := InitToDo(a.DB)
	a.Router.HandleFunc("/api/todos", postAPI.FindAllPosts()).Methods("GET")
	a.Router.HandleFunc("/api/todos", postAPI.CreateToDo()).Methods("POST")
	a.Router.HandleFunc("/api/todos/{id:[0-9]+}", postAPI.FindByID()).Methods("GET")
	a.Router.HandleFunc("/api/todos/{id:[0-9]+}", postAPI.DeleteToDo()).Methods("DELETE")
}

func InitToDo(db *gorm.DB) api.ToDoAPI {
	todoRepository := repository.NewRepository(db)
	todoService := service.NewToDoService(todoRepository)
	todoAPI := api.NewToDoAPI(todoService)
	todoAPI.Migrate()
	return todoAPI
}
