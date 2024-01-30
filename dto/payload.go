package dto

type AddTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TodoResponse struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"created_at"`
}
