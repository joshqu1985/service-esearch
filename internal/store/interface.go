package store

//go:generate mockgen -destination=../mock/store_mock.go -package=mock github.com/joshqu1985/service-esearch/internal/store Store

import (
	"github.com/joshqu1985/service-esearch/internal/model"
)

type Store interface {
	UsersByName(name string, offset, limit int) ([]model.SearchInfo, error)
	UsersByNear(lat, lon float64, gender, offset, limit int) ([]model.SearchInfo, error)
}
