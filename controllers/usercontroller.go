package controllers
import (
  "jwt-authentication-golang/database"
  "jwt-authentication-golang/models"
  "net/http"
  "github.com/gin-gonic/gin"
)
func RegisterUser(context *gin.Context) {
  var user models.User
  // Add data from client to User object
  if err := context.ShouldBindJSON(&user); err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  if err := user.HashPassword(user.Password); err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  // Store User data in database
  record := database.Instance.Create(&user)
  if record.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
    context.Abort()
    return
  }
  // Send client success response
  context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}