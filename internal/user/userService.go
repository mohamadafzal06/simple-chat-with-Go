package user

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamadafzal06/simple-chat/internal/util"
)

const (
	secretKey = "secret"
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

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserResp, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserResp{}, fmt.Errorf("cannot get user: %v\n", err)
	}

	err = util.CheckHashPassword(u.Password, req.Password)
	if err != nil {
		return &LoginUserResp{}, fmt.Errorf("wrong password -> authentication failed: %v\n", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserResp{}, fmt.Errorf("error while signing token: %v\n", err)
	}

	return &LoginUserResp{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}
