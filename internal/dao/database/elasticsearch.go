package database

import (
	"github.com/joshqu1985/fireman/pkg/store/elasticsearch"
)

type db struct {
	Pool *elasticsearch.Pool
}

func NewRepository(pool *elasticsearch.Pool) DB {
	return &db{Pool: pool}
}
