package biz_context

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
)

type User struct {
	UserID   string
	UserName string
}

type BizContext struct {
	context.Context
	ProjectCode string `header:"projectCode"`
	Lang        string `header:"lang"`
	User
	Resources []string
}

func NewBizContext(ctx context.Context, c *app.RequestContext) (*BizContext, error) {
	userID, ok := c.Get("user_id")
	if !ok {
		return nil, errors.New("user id not found")
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return nil, errors.New("user id is not string")
	}

	userName, ok := c.Get("user_name")
	if !ok {
		return nil, errors.New("user name not found")
	}

	userNameStr, ok := userName.(string)
	if !ok {
		return nil, errors.New("user name is not string")
	}

	ret := &BizContext{
		Context: ctx,
		User: User{
			UserID:   userIDStr,
			UserName: userNameStr,
		},
	}

	resources, ok := c.Get("resources")
	if ok {
		ret.Resources, ok = resources.([]string)
		if !ok {
			return nil, errors.New("resources is not []string")
		}
	}

	// this will bind the struct which has the tag "header" from the request header
	err := c.BindHeader(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func Background() *BizContext {
	return &BizContext{
		Context: context.Background(),
	}
}
