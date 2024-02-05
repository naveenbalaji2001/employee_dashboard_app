package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naveenbalaji2001/employee_dashboard_app/config"
	"github.com/naveenbalaji2001/employee_dashboard_app/routes"

	_ "github.com/go-redis/redis/v8"
	_ "gorm.io/gorm"
)

func main() {

	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}

	redisClient := config.SetupRedis()

	recentDelClient := config.SetupRecentDeletedRedis()

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router, db, redisClient, recentDelClient)
	router.Run(":8080")
}
