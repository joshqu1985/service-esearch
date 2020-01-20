package database

import (
	"github.com/joshqu1985/fireman/pkg/store/es"
)

type db struct {
	Pool *es.Pool
}

func NewRepository(pool *es.Pool) DB {
	return &db{Pool: pool}
}
