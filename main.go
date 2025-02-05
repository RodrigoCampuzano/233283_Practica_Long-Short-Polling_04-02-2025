package main

import (
    "Practica/src/infraestrcuture/db"
    "Practica/src/infraestrcuture/controllers"
    "Practica/src/domain/repositories"
    "Practica/src/infraestrcutureroutes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializar la base de datos
    db := db.ConnectionDB()
    defer db.Close()

    // Repositorio
    userRepo := repositories.NewUserRepository(db)

    // Controlador
    userController := controllers.NewUserController(userRepo)

    // Configurar el enrutador
    router := gin.Default()
    routes.SetupUserRoutes(router, userController)

    // Iniciar el servidor
    router.Run(":8080")
}