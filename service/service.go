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

type UpdateTodoRequest struct {
	Id          int64
	Title       string
	Description string
	Done        bool
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

func (s Service) GetAll() ([]TodoResponse, error) {
	var resp []TodoResponse
	tasks, err := db.New(s.dbCfg).GetAll()

	for _, task := range tasks {
		response := TodoResponse{
			Id:          task.Id,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
			CreatedAt:   task.CreatedAt.String(),
		}
		resp = append(resp, response)
	}

	return resp, err
}

// todo handle if id not found
func (s Service) Update(req UpdateTodoRequest) error {
	task := model.Task{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
	}

	return db.New(s.dbCfg).Update(task)
}

func (s Service) DeleteByID(id int64) error {
	return db.New(s.dbCfg).DeleteByID(id)
}
