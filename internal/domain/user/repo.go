package user

import (
	"github.com/morris-zheng/go-slim-micro-usersvc/internal/domain"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(svc *domain.ServiceContext) *Repo {
	return &Repo{
		DB: svc.DB,
	}
}
