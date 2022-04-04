package resolver

import (
	"context"
)

func (r *Resolver) GetHealthCheck(context.Context) string {
	return "Ok"
}
