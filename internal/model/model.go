package model

type UsersByNameArgs struct {
	Name   string `json:"name"`
	Offset int64  `json:"offset"`
	Limit  int    `json:"limit"`
}

type UsersByNearArgs struct {
	Lon    float64 `json:"lon"`
	Lat    float64 `json:"lat"`
	Offset int64   `json:"offset"`
	Limit  int     `json:"limit"`
}

type SearchInfo struct {
	Source   []byte
	Distance float64
}

type SearchInfos []SearchInfo
