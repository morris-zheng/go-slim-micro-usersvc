package user

import (
	"context"
	"log"

	"github.com/morris-zheng/go-slim-micro-usersvc/export/user"
	"github.com/morris-zheng/go-slim-micro-usersvc/export/usersvc"
	"github.com/morris-zheng/go-slim-micro-usersvc/internal/domain"
	userDomain "github.com/morris-zheng/go-slim-micro-usersvc/internal/domain/user"

	"github.com/jinzhu/copier"
)

type Service struct {
	user.UnimplementedServiceServer
	svc *domain.ServiceContext
	uc  *userDomain.UseCase
}

func NewService(svc *domain.ServiceContext) *Service {
	return &Service{
		svc: svc,
		uc:  userDomain.NewUseCase(svc),
	}
}

func (s *Service) Create(ctx context.Context, req *user.User) (*usersvc.OpResp, error) {
	o := userDomain.User{}
	err := copier.Copy(&o, req)
	if err != nil {
		return nil, err
	}

	err = s.uc.Create(ctx, &o)
	if err != nil {
		return nil, err
	}

	return &usersvc.OpResp{
		Msg:  "success",
		Code: 0,
	}, nil
}

func (s *Service) Get(ctx context.Context, req *user.UserById) (*user.User, error) {
	log.Println("get")
	return &user.User{
		Id:         req.Id,
		CreateTime: "2",
		UpdateTime: "3",
	}, nil
}
