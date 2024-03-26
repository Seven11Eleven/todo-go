package routes

import (
	"net/http"
	"to-do/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos", controllers.GetTodo)
	r.GET("/todos/:id", controllers.GetTasks)
	r.POST("/todos/:id/tasks", controllers.CreateTask)
	r.DELETE("/todos/:id", controllers.DeleteTodoList)
	r.PUT("todos/:id/tasks/:taskid", controllers.SetCompleteOrIncomplete)
	return r
}

