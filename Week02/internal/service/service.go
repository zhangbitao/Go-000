package service

import (
	"os"

	"Week02/internal/biz"

	"github.com/pkg/errors"
)

type Service interface {
	GetUserName() (name string, err error)
}

type service struct {
	biz biz.Biz
}

func New(bz biz.Biz) Service {
	return &service{biz: bz}
}

func (s *service) GetUserName() (name string, err error) {
	if len(os.Args) != 2 {
		return "", errors.New("invalid argument")
	}
	return s.biz.GetUserName(os.Args[1])
}
