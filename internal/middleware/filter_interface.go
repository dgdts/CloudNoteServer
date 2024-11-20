package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type Filter interface {
	Init()
	GetOrder() int
	DoFilter(ctx context.Context, c *app.RequestContext)
}

func GetAllFilters() []Filter {
	filters := []Filter{
		&FilterValidate{},
		&FilterLogin{},
		&FilterResource{},
	}

	for _, filter := range filters {
		filter.Init()
	}

	return filters
}
