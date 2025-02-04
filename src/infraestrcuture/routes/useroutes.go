package routes

import (
    "github.com/gin-gonic/gin"
    "Practica/src/infraestrcuture/controllers"
)

func SetupUserRoutes(router *gin.Engine, userController *controllers.UserController) {
    router.GET("/users/:id", userController.GetUserByID)
    router.GET("/users/short-polling", userController.ShortPolling)
}