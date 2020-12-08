package biz

import (
	"strconv"

	"Week02/internal/code"
	"Week02/internal/dao"

	"github.com/pkg/errors"
)

type Biz interface {
	GetUserName(id string) (name string, err error)
}

type biz struct {
	dao dao.Dao
}

func New(dao dao.Dao) Biz {
	return &biz{dao: dao}
}

func (b *biz) GetUserName(id string) (name string, err error) {
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse id")
	}

	name, err = b.dao.GetUserName(userID)
	if err == code.ErrRecordNotFound {
		name = "Unknown"
		err = nil
	}
	return
}
