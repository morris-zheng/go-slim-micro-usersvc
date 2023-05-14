package user

import (
	"context"
	"time"

	"github.com/morris-zheng/go-slim-micro-usersvc/internal/domain"
)

type UseCase struct {
	svc  *domain.ServiceContext
	repo *Repo
}

func NewUseCase(svc *domain.ServiceContext) *UseCase {
	return &UseCase{
		svc:  svc,
		repo: NewRepo(svc),
	}
}

func (uc *UseCase) Create(ctx context.Context, u *User) error {
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()

	err := uc.repo.DB.Create(u).Error
	if err != nil {
		return err
	}

	return nil
}
