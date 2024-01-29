package service

import (
	"TodoBackend/db"
	"TodoBackend/model"
	"fmt"
	"strings"
)

type AddTodoRequest struct {
	Title       string
	Description string
}

type TodoResponse struct {
	Id          int64
	Title       string
	Description string
	Done        bool
	CreatedAt   string
}

type Service struct {
	dbCfg db.Config
}

func New(dbConfig db.Config) Service {
	return Service{dbCfg: dbConfig}
}

func (s Service) Add(req AddTodoRequest) (TodoResponse, error) {
	if len(strings.TrimSpace(req.Title)) == 0 || len(strings.TrimSpace(req.Description)) == 0 {
		return TodoResponse{}, fmt.Errorf("invalid input")
	}

	task := model.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	savedTodo, _ := db.New(s.dbCfg).Save(task)

	return TodoResponse{
		Id:          savedTodo.Id,
		Title:       savedTodo.Title,
		Description: savedTodo.Description,
		Done:        savedTodo.Done,
		CreatedAt:   savedTodo.CreatedAt.String(),
	}, nil

}

func (s Service) GetById(id int64) (TodoResponse, error) {
	task, err := db.New(s.dbCfg).GetByID(id)
	if err != nil {
		return TodoResponse{}, err
	}

	return TodoResponse{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
		CreatedAt:   task.CreatedAt.String(),
	}, nil

}

/*

func getAll() []TodoResponse {

}*/
