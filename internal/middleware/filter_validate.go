package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgdts/CloudNoteServer/internal/response"
)

const (
	FilterValidateName = "filter_validate"
)

type FilterValidate struct {
	excludePaths []string
}

var _ Filter = (*FilterValidate)(nil)

func (f *FilterValidate) Name() string {
	return FilterValidateName
}

func (f *FilterValidate) Init() {
	excludePaths := []string{
		"/ping",
		"/api/v1/auth/login",
		"/api/v1/auth/register",
	}

	f.excludePaths = excludePaths
}

func (f *FilterValidate) GetOrder() int {
	return 20
}

func (f *FilterValidate) DoFilter(ctx context.Context, c *app.RequestContext) {
	err := f.doFilter(ctx, c)
	if err != nil {
		hlog.CtxErrorf(ctx, "filter validate failed: %v", err)
		response.JSONError(c, err)
		c.Abort()
		return
	}
	c.Next(ctx)
}

func (f *FilterValidate) doFilter(ctx context.Context, c *app.RequestContext) error {
	path := string(c.Path())
	for _, excludePath := range f.excludePaths {
		if strings.HasPrefix(path, excludePath) {
			return nil
		}
	}

	bizCtx := getBizContext(c)

	if bizCtx == nil {
		err := errors.New("biz context not found")
		hlog.CtxErrorf(ctx, "%v", err)
		return err
	}

	return nil
}
