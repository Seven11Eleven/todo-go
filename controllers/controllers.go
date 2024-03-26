package controllers

import (
	"fmt"
	"context"
	// "encoding/json"
	"net/http"
	"strconv"
	"to-do-listik/models"

	"github.com/gin-gonic/gin"
	
	."to-do-listik/database"
)

func CreateTodo(c *gin.Context){
	var todoList models.TodoList
	if err := c.BindJSON(&todoList); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Conn.QueryRow(context.Background(), "INSERT INTO todo_lists (title, description) VALUES ($1, $2) RETURNING id", todoList.Title, todoList.Description).Scan(&todoList.ID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todoList)
}

func CreateTask(c *gin.Context) {
	todolistID, _ := strconv.Atoi(c.Param("id"))
	var task models.Task 
	if err := c.BindJSON(&task); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.TodoListID = todolistID
	err := Conn.QueryRow(context.Background(), "INSERT INTO tasks (title, description, completed, todo_list_id) VALUES ($1, $2, $3, $4) RETURNING id", task.Title, task.Description, task.Completed, task.TodoListID).Scan(&task.ID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func GetTodo(c *gin.Context){
	rows, err := Conn.Query(context.Background(), "SELECT * FROM todo_lists")
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	
	todoLists := []models.TodoList{}
	for rows.Next() {
		var todoList models.TodoList
		err := rows.Scan(&todoList.ID, &todoList.Title, &todoList.Description)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todoLists = append(todoLists, todoList)


	}
	c.JSON(http.StatusOK, todoLists)
}

func GetTasks(c *gin.Context){
	todoListID, _ := strconv.Atoi(c.Param("id"))
	rows, err := Conn.Query(context.Background(), "SELECT * FROM tasks WHERE todo_list_id = $1 ORDER BY id ASC", todoListID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.TodoListID)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}

func DeleteTodoList(c *gin.Context){
	todoListID, _ := strconv.Atoi(c.Param("id"))

	
	 tasksExist, err := CheckTasksExist(todoListID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check tasks existence"})
        return
    }
    if tasksExist {
        if err := DeleteAllTasksByToDoID(todoListID); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
    }
		
	
	result, err := Conn.Exec(context.Background(), "DELETE FROM todo_lists WHERE id = $1", todoListID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	rowAffect := result.RowsAffected()

	if rowAffect == 0{
		c.JSON(http.StatusNotFound, gin.H{"error": "To-do List not found, man"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("To-do list with ID %d deleted successfully!", todoListID)})
}

func DeleteAllTasksByToDoID(ListID int) error{
	_, err := Conn.Exec(context.Background(),"DELETE FROM tasks WHERE todo_list_id = $1", ListID)
	return err
	}

func CheckTasksExist(todoListID int) (bool, error) {
	var tasksExist bool
	err := Conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM tasks WHERE todo_list_id = $1)", todoListID).Scan(&tasksExist)
	if err != nil {
		return false, err
	}
	return tasksExist, nil
	}

func SetCompleteOrIncomplete(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("taskid"))
	_, err := Conn.Exec(context.Background(), "UPDATE tasks SET completed = NOT completed WHERE id = $1", taskID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус выполнения задачи с ID " + strconv.Itoa(taskID) + " успешно обновлен"})

}