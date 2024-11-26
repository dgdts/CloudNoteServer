package response

import (
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

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
