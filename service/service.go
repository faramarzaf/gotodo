package service

import (
	"TodoBackend/db"
	"TodoBackend/dto"
	"TodoBackend/model"
	"fmt"
	"strings"
)

type Service struct {
	dbCfg db.Config
}

func New(dbConfig db.Config) Service {
	return Service{dbCfg: dbConfig}
}

func (s Service) Add(req dto.AddTodoRequest) (dto.TodoResponse, error) {
	if len(strings.TrimSpace(req.Title)) == 0 || len(strings.TrimSpace(req.Description)) == 0 {
		return dto.TodoResponse{}, fmt.Errorf("invalid input")
	}

	task := model.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	savedTodo, _ := db.New(s.dbCfg).Save(task)

	return dto.TodoResponse{
		Id:          savedTodo.Id,
		Title:       savedTodo.Title,
		Description: savedTodo.Description,
		Done:        savedTodo.Done,
		CreatedAt:   savedTodo.CreatedAt.String(),
	}, nil

}

func (s Service) GetById(id int64) (dto.TodoResponse, error) {
	task, err := db.New(s.dbCfg).GetByID(id)
	if err != nil {
		return dto.TodoResponse{}, err
	}

	return dto.TodoResponse{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
		CreatedAt:   task.CreatedAt.String(),
	}, nil

}

func (s Service) GetAll() ([]dto.TodoResponse, error) {
	var resp []dto.TodoResponse
	tasks, err := db.New(s.dbCfg).GetAll()

	for _, task := range tasks {
		response := dto.TodoResponse{
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
func (s Service) Update(req dto.UpdateTodoRequest) error {
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
