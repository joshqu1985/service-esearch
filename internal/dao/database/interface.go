package database

//go:generate mockgen -destination=../mock/store_mock.go -package=mock github.com/joshqu1985/service-esearch/internal/dao/database DB

import (
	"context"

	"github.com/joshqu1985/service-esearch/internal/model"
)

type DB interface {
	UsersByName(ctx context.Context, args *model.UsersByNameArgs) (model.SearchInfos, error)
	UsersByNear(ctx context.Context, args *model.UsersByNearArgs) (model.SearchInfos, error)
}
