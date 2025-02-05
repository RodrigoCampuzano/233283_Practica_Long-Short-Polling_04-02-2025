package application

import (
    "Practica/src/domain/entities"
    "Practica/src/domain/repositories"
)

type GetUserByIDUseCase struct {
    userRepo repositories.UserRepository
}

func NewGetUserByIDUseCase(userRepo repositories.UserRepository) *GetUserByIDUseCase {
    return &GetUserByIDUseCase{userRepo: userRepo}
}

func (uc *GetUserByIDUseCase) Execute(id int) (*entities.User, error) {
    return uc.userRepo.GetUserByID(id)
}