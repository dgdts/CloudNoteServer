package middleware

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgdts/CloudNoteServer/biz/biz_context"
	"github.com/dgdts/CloudNoteServer/internal/response"
)

const (
	FilterResourceName = "filter_resource"
)

type FilterResource struct {
	excludePaths []string
}

var _ Filter = &FilterResource{}

func (f *FilterResource) Name() string {
	return FilterResourceName
}

func (f *FilterResource) Init() {
	excludePaths := []string{
		"/ping",
		"/api/v1/auth/login",
		"/api/v1/auth/register",
	}
	f.excludePaths = excludePaths

}

func (f *FilterResource) GetOrder() int {
	return 10
}

func (f *FilterResource) DoFilter(ctx context.Context, c *app.RequestContext) {
	err := f.doFilter(ctx, c)
	if err != nil {
		hlog.CtxErrorf(ctx, "%v", err)
		response.JSONError(c, err)
		c.Abort()
		return
	}
	c.Next(ctx)
}

func (f *FilterResource) doFilter(ctx context.Context, c *app.RequestContext) error {
	path := string(c.Path())
	for _, excludePath := range f.excludePaths {
		if strings.HasPrefix(path, excludePath) {
			return nil
		}
	}

	bizCtx, err := biz_context.NewBizContext(ctx, c)
	if err != nil {
		return err
	}

	iamResourceByte := c.GetHeader("X-Iam-Resource")
	iamResources := make([]string, 0)

	if len(iamResourceByte) > 0 {
		iamResources = strings.Split(string(iamResourceByte), ",")
		c.Set("resources", iamResources)
	}

	hlog.CtxInfof(ctx, "iam resources: %v", iamResources)

	c.Set("biz_context", bizCtx)

	return nil
}
