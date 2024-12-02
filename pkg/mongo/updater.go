package mongo

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/builder/update"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Update struct {
	collection  *mongoDriver.Collection
	opts        []*mongoOptions.UpdateOptions
	replaceOpts []*mongoOptions.ReplaceOptions

	u *update.Builder
	q *query.Builder
}

func Updater(collection string, opts ...CollectionOption) *Update {
	ret := &Update{}
	ret.q = query.NewBuilder()
	ret.u = update.NewBuilder()
	ret.collection = RawCollection(collection, opts...)
	return ret
}

func (u *Update) WithUpdateOptions(opts ...*mongoOptions.UpdateOptions) *Update {
	u.opts = append(u.opts, opts...)
	return u
}

func (u *Update) WithEqFilter(pairs ...any) *Update {
	appendEq(u.q, pairs...)
	return u
}

func (u *Update) UpdateOne(ctx context.Context, fieldValues ...any) updateResultInterface {
	if len(fieldValues) > 0 {
		var k string
		for i, pair := range fieldValues {
			if i%2 == 0 {
				k = pair.(string)
			} else {
				u.u.Set(k, pair)
			}
		}
	}

	r, err := u.collection.UpdateOne(ctx, u.q.Build(), u.u.Build(), u.opts...)
	if err != nil {
		return newErrorResult(err)
	}
	return newUpdateResult(r)
}

func (u *Update) WithReplaceOptions(opts ...*mongoOptions.ReplaceOptions) *Update {
	u.replaceOpts = append(u.replaceOpts, opts...)
	return u
}

func (u *Update) ReplaceOne(ctx context.Context, data any) updateResultInterface {
	r, err := u.collection.ReplaceOne(ctx, u.q.Build(), data, u.replaceOpts...)
	if err != nil {
		return newErrorResult(err)
	}
	return newUpdateResult(r)
}

func appendEq(q *query.Builder, pairs ...any) {
	var k string

	for i, pair := range pairs {
		if i%2 == 0 {
			k = pair.(string)
		} else {
			q.Eq(k, pair)
		}
	}
}
