package services

import (
	"log"

	"github.com/marufnwu/HexaArchProject/domain"
)

type MyServiceInterface interface{
	 GetUsers() []domain.User
}

type MyService struct{
	repo domain.MyRepository
}

func (service MyService) GetUsers() []domain.User {
	users, err := service.repo.FindAllUser()
	if err!=nil{
		log.Fatal(err.Error())
	}
	return users
}

func NewService() MyServiceInterface{
	return MyService{
		repo: domain.NewRepository(),
	}
}