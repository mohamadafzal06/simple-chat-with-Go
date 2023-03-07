package user

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/mohamadafzal06/simple-chat/internal/util"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(r Repository) *service {
	return &service{
		r,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserResp, error) {
	c, cancle := context.WithTimeout(ctx, s.timeout)
	defer cancle()

	hashpass, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("cannot hash the password: %v\n", err)
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashpass,
	}

	r, err := s.Repository.CreateUser(c, u)
	if err != nil {
		return nil, fmt.Errorf("cannot create user: %v\n", err)
	}

	rUser := &CreateUserResp{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return rUser, nil
}
