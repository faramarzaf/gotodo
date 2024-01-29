package dto

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
