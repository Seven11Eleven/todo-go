package models

type TodoList struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
    TodoListID  int    `json:"todo_list_id"`
}
