package config

import (
  "github.com/gin-gonic/gin"

  "github.com/augusto-queirantes/learning_go/controllers"
)

func Server() {
  router := gin.Default()

  router.GET("/", controllers.HomeIndex)
  router.GET("/message/:message", controllers.ShowMessage)

  router.Run(":3000")
}
