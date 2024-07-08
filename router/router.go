package routes

import (
	"appinsights/test/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudentByID)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	return r
}
