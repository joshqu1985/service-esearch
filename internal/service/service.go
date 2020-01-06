package service

import (
	"github.com/joshqu1985/service-esearch/internal/dao/database"
)

type Service struct {
	DB database.DB
}

func New(db database.DB) *Service {
	return &Service{db}
}
