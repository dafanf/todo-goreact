// routes/task_routes.go

package routes

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

// Set up the router and wire routes to controller methods.
func SetupTaskRoutes(router *gin.Engine, taskController *controller.TaskController) {
	// Define routes and bind them to controller methods.
	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	// Define other routes...
}
