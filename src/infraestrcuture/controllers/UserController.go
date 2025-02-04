package controllers

import (
    "net/http"
    "strconv"
    "Practica/src/domain/entities"
    "Practica/src/domain/repositories"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    userRepo repositories.UserRepository
}

func NewUserController(userRepo repositories.UserRepository) *UserController {
    return &UserController{userRepo: userRepo}
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    user, err := c.userRepo.GetUserByID(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) ShortPolling(ctx *gin.Context) {
    users, err := c.userRepo.CheckForChanges()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, users)
}