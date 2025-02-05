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
	FilterLoginName = "filter_login"
)

type FilterLogin struct {
	excludePaths []string
}

var _ Filter = (*FilterLogin)(nil)

func (f *FilterLogin) Name() string {
	return FilterLoginName
}

func (f *FilterLogin) Init() {
	excludePaths := []string{
		"/ping",
		"/api/v1/auth/login",
		"/api/v1/auth/register",
	}

	f.excludePaths = excludePaths
}

func (f *FilterLogin) GetOrder() int {
	return 0
}

func (f *FilterLogin) DoFilter(ctx context.Context, c *app.RequestContext) {
	err := f.doFilter(ctx, c)
	if err != nil {
		hlog.CtxErrorf(ctx, "filter login failed: %v", err)
		response.JSONError(c, err)
		c.Abort()
		return
	}
	c.Next(ctx)
}

func (f *FilterLogin) doFilter(ctx context.Context, c *app.RequestContext) error {
	path := string(c.Path())
	for _, excludePath := range f.excludePaths {
		if strings.HasPrefix(path, excludePath) {
			return nil
		}
	}

	auth := c.GetHeader("Authorization")
	if len(auth) == 0 {
		return errors.New("auth is empty")
	}

	claims, err := getAuthFromJWTToken(string(auth))
	if err != nil {
		hlog.CtxErrorf(ctx, "get auth from jwt token failed: %v", err)
		return err
	}

	c.Set(UserNameKey, claims[UserNameKey])
	c.Set(UserIDKey, claims[UserIDKey])

	return nil
}
