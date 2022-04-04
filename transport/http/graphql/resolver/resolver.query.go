package resolver

import "context"

type QueryResolver struct {
	Field string
}

func (r *Resolver) Query(ctx context.Context) *QueryResolver {
	return &QueryResolver{Field: "ok"}
}
