package handler

import (
	"context"
	"encoding/json"

	"github.com/joshqu1985/fireman/util/page"
	"github.com/joshqu1985/protos/pb"

	"github.com/joshqu1985/service-esearch/internal/model"
)

func (h *Handler) UsersByName(ctx context.Context, in *pb.UsersByNameArgs) (*pb.UserInfos, error) {
	resp := &pb.UserInfos{}

	ptok := &page.Token{}
	if err := ptok.Decode(in.PageToken); err != nil {
		return resp, err
	}

	rows, err := h.Service.UsersByName(ctx, &model.UsersByNameArgs{
		Name:   in.Name,
		Offset: ptok.Offset,
		Limit:  ptok.Limit,
	})
	if err != nil || len(rows) == 0 {
		return resp, err
	}

	return resp, h.packUsers(resp, rows, ptok)
}

func (h *Handler) UsersByNear(ctx context.Context, in *pb.UsersByNearArgs) (*pb.UserInfos, error) {
	resp := &pb.UserInfos{}

	ptok := &page.Token{}
	if err := ptok.Decode(in.PageToken); err != nil {
		return resp, err
	}

	rows, err := h.Service.UsersByNear(ctx, &model.UsersByNearArgs{
		Lat:    in.Lat,
		Lon:    in.Lon,
		Offset: ptok.Offset,
		Limit:  ptok.Limit,
	})
	if err != nil || len(rows) == 0 {
		return resp, err
	}

	return resp, h.packUsers(resp, rows, ptok)
}

func (h *Handler) packUsers(resp *pb.UserInfos, rows model.SearchInfos, ptok *page.Token) error {
	for _, row := range rows {
		user := pb.UserInfo{}
		if err := json.Unmarshal(row.Source, &user); err != nil {
			continue
		}
		resp.Items = append(resp.Items, &user)
	}

	ptok.Offset += int64(len(rows))
	resp.PageToken = ptok.Encode()
	return nil
}
