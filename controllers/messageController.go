package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func ShowMessage(context *gin.Context) {
  message := context.Param("message")

  context.JSON(http.StatusOK, gin.H{"message": message})
}
