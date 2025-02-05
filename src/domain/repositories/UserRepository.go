package repositories

import "Practica/src/domain/entities"

type UserRepository interface {
    GetUserByID(id int) (*entities.User, error)
    UpdateUser(user *entities.User) error
    CheckForChanges() ([]entities.User, error) 
    WaitForChanges() ([]entities.User, error)  
}