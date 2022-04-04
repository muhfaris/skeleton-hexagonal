package rslvmodel

import "github.com/graph-gophers/graphql-go"

type AccountDataResolver struct {
	StatusRslv int
	DataRslv   AccountResolver
}

func (ad *AccountDataResolver) ID() graphql.ID {
	return graphql.ID(ad.DataRslv.Id.String())
}

func (ad *AccountDataResolver) Status() int32 {
	return int32(ad.StatusRslv)
}

func (ad *AccountDataResolver) Data() *AccountResolver {
	return &ad.DataRslv
}
