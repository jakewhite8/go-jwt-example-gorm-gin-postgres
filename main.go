package main
import (
  "jwt-authentication-golang/controllers"
  "jwt-authentication-golang/database"
  "jwt-authentication-golang/middleware"
  "github.com/gin-gonic/gin"
)
func main() {
  // Initialize Database
  database.Connect()
  database.Migrate()
  // Initialize Router
  router := initRouter()
  router.Run(":8080")
}
func initRouter() *gin.Engine {
  router := gin.Default()
  api := router.Group("/api")
  {
    api.POST("/token", controllers.GenerateToken)
    api.POST("/user/register", controllers.RegisterUser)
    secured := api.Group("/secured").Use(middleware.Auth())
    {
      secured.GET("/ping", controllers.Ping)
    }
  }
  return router
}