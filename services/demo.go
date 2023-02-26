package services

import (
	"labs/repository"

	"gorm.io/gorm"
)

type DemoService interface {
	WithTrx(*gorm.DB) demoService
}

type demoService struct {
	usersRepository repository.UsersRepository
}

func NewDemoService(r repository.UsersRepository) demoService {
	return demoService{
		usersRepository: r,
	}
}

func (g demoService) WithTrx(trxHandle *gorm.DB) demoService {
	g.usersRepository = g.usersRepository.WithTrx(trxHandle)
	return g
}
