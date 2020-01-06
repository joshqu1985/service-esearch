package service

import (
	"context"

	"github.com/joshqu1985/service-esearch/internal/model"
)

func (s *Service) UsersByName(ctx context.Context, args *model.UsersByNameArgs) (model.SearchInfos, error) {
	return s.DB.UsersByName(ctx, args)
}

func (s *Service) UsersByNear(ctx context.Context, args *model.UsersByNearArgs) (model.SearchInfos, error) {
	return s.DB.UsersByNear(ctx, args)
}
