package repository

import (
	"log"

	"gorm.io/gorm"
)

type UsersRepository interface {
	WithTrx(*gorm.DB) usersRepository
}

type usersRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) usersRepository {
	return usersRepository{
		DB: db,
	}
}

func (g usersRepository) WithTrx(trxHandle *gorm.DB) usersRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return g
	}
	g.DB = trxHandle
	return g
}
