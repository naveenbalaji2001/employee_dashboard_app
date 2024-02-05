package routes

import (
	// "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/naveenbalaji2001/employee_dashboard_app/controller"
	"github.com/naveenbalaji2001/employee_dashboard_app/dal"
	"github.com/naveenbalaji2001/employee_dashboard_app/service"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client, recentDelClient *redis.Client) {
	todoDAL := dal.NewTodoDAL(db, redisClient, recentDelClient, "recent_del:", time.Minute*5)

	todoService := service.NewTodoService(todoDAL)

	todoController := controller.NewTodoController(todoService)

	// router.Use(func(ctx *gin.Context) {
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization")
	// 	if ctx.Request.Method == "OPTIONS" {
	// 		ctx.AbortWithStatus(http.StatusNoContent)
	// 	  return
	// 	}
	// 	ctx.Next()
	//   })
	router.GET("/login", todoController.Login)
	router.GET("/callback", todoController.Callback)
	router.GET("/user", todoController.GetUserHandler)

	router.POST("/register", todoController.RegisterUser)
	router.GET("/getregister", todoController.GetRegisterUser)
	router.POST("/login", todoController.VerifyLogin)

	router.POST("/todos", todoController.CreateTodo)
	router.GET("/todos/:id", todoController.GetTodo)
	router.GET("/todos", todoController.GetTodos)
	router.PUT("/todos/:id", todoController.UpdateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	router.GET("/recent_del", todoController.GetRecentDeletedTodos)
	router.POST("/clear", todoController.ClearData)
}
