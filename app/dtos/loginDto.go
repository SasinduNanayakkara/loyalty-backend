package dtos

import "github.com/sasinduNanayakkara/loyalty-backend/app/models"

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponseDto struct {
	Token    string          `json:"token"`
	Customer models.Customer `json:"customer"`
}
