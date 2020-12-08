package dao

import (
	"database/sql"

	"Week02/internal/code"

	"github.com/pkg/errors"
)

type Dao interface {
	GetUserName(id int64) (name string, err error)
}

type dao struct {
	db *sql.DB
}

func New(db *sql.DB) Dao {
	return &dao{db: db}
}

func (d *dao) GetUserName(id int64) (name string, err error) {
	rows, err := d.db.Query("select name from user where id = ? limit 1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", code.ErrRecordNotFound
		}

		return "", errors.Wrap(err, "select failed")
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			continue
		}
	}

	return
}
