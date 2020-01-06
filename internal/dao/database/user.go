package database

import (
	"context"

	"github.com/olivere/elastic"

	"github.com/joshqu1985/service-esearch/internal/model"
)

func (d *db) UsersByName(ctx context.Context, args *model.UsersByNameArgs) (model.SearchInfos, error) {
	items := make(model.SearchInfos, 0)
	if int(args.Offset)+args.Limit > 10000 {
		return items, nil
	}

	handle := func(c *elastic.Client) error {
		queryer := elastic.NewBoolQuery().
			Should(elastic.NewQueryStringQuery(args.Name).Field("name"))
		res, err := c.Search().Index("user").Query(queryer).
			From(int(args.Offset)).Size(args.Limit).Do(context.Background())
		if err != nil {
			return err
		}
		if res.Hits.TotalHits == 0 {
			return nil
		}

		for _, hit := range res.Hits.Hits {
			items = append(items, model.SearchInfo{Source: *hit.Source})
		}
		return nil
	}
	return items, d.Pool.Doit(ctx, handle)
}

func (d *db) UsersByNear(ctx context.Context, args *model.UsersByNearArgs) (model.SearchInfos, error) {
	items := make(model.SearchInfos, 0)
	if int(args.Offset)+args.Limit > 10000 {
		return items, nil
	}

	handle := func(c *elastic.Client) error {
		queryer := elastic.NewBoolQuery().Filter(elastic.NewTermQuery("state", 0),
			elastic.NewGeoDistanceQuery("geo").Point(args.Lat, args.Lon).Distance("10km"))
		sorter := elastic.NewGeoDistanceSort("geo").
			Point(args.Lat, args.Lon).Asc().Unit("km").SortMode("min").GeoDistance("plane")
		res, err := c.Search().Index("user").Query(queryer).
			SortBy(sorter).From(int(args.Offset)).Size(args.Limit).Do(context.Background())
		if err != nil {
			return err
		}
		if res.Hits.TotalHits == 0 {
			return nil
		}

		for _, hit := range res.Hits.Hits {
			item := model.SearchInfo{Source: *hit.Source}
			if len(hit.Sort) != 0 {
				item.Distance = hit.Sort[0].(float64)
			}
			items = append(items, item)
		}
		return nil
	}
	return items, d.Pool.Doit(ctx, handle)
}
