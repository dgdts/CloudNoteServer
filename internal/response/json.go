package response

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgdts/CloudNoteServer/biz/biz_context"
)

type JSONHandler[Req any, Res any] func(c *biz_context.BizContext, req *Req) (*Res, error)

func JSONError(c *app.RequestContext, err error) {
	resp := NewResultFromError(err)

	// bizStatus used to monitoring instrument
	c.Response.Header.Set("bizStatus", strconv.Itoa(resp.Status))
	c.JSON(http.StatusOK, resp)
}

func JSONSuccess(c *app.RequestContext, data interface{}) {
	resp := NewResultWithData(data)
	c.JSON(http.StatusOK, resp)
}

func JSON[Req any, Res any](ctx context.Context, c *app.RequestContext, handler JSONHandler[Req, Res]) {
	var req Req
	err := c.BindAndValidate(&req)
	if err != nil {
		JSONError(c, err)
		return
	}

	bizCtx, err := biz_context.NewBizContext(ctx, c)
	if err != nil {
		JSONError(c, err)
		return
	}

	resp, err := handler(bizCtx, &req)
	if err != nil {
		JSONError(c, err)
		return
	}

	JSONSuccess(c, resp)
}
