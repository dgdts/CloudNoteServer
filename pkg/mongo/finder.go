package mongo

import (
	"context"
	"errors"
	"regexp"

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Find struct {
	collection   *mongoDriver.Collection
	opts         []*mongoOptions.FindOptions
	oneOpts      []*mongoOptions.FindOneOptions
	replaceOpts  []*mongoOptions.FindOneAndReplaceOptions
	updateOpts   []*mongoOptions.FindOneAndUpdateOptions
	countOpts    []*mongoOptions.CountOptions
	orConditions []bson.D

	sorts  []interface{}
	fields []string
	q      *query.Builder
}

func Finder(collection string, opts ...CollectionOption) *Find {
	ret := &Find{
		collection: RawCollection(collection, opts...),
	}
	ret.q = query.NewBuilder()

	return ret
}

func (f *Find) WithSort(sortPairs ...interface{}) *Find {
	f.sorts = append(f.sorts, sortPairs...)
	return f
}

func (f *Find) In(key string, values ...any) *Find {
	f.q.In(key, values...)
	return f
}

func (f *Find) NotIn(key string, values ...any) *Find {
	f.q.Nin(key, values...)
	return f
}

func (f *Find) Gt(key string, value any) *Find {
	f.q.Gt(key, value)
	return f
}

func (f *Find) Gte(key string, value any) *Find {
	f.q.Gte(key, value)
	return f
}

func (f *Find) Lt(key string, value any) *Find {
	f.q.Lt(key, value)
	return f
}

func (f *Find) Lte(key string, value any) *Find {
	f.q.Lte(key, value)
	return f
}

func (f *Find) Ors(fieldValues ...any) *Find {
	var field string
	condictions := []bson.D{}
	for i, fieldValue := range fieldValues {
		if i%2 == 0 {
			field = fieldValue.(string)
		} else {
			condictions = append(condictions, query.Eq(field, fieldValue))
		}
	}
	f.orConditions = append(f.orConditions, condictions...)
	return f
}

func (f *Find) Or(conditions ...bson.D) *Find {
	f.orConditions = append(f.orConditions, conditions...)
	return f
}

func (f *Find) Between(key string, start, end any) *Find {
	f.q.Gte(key, start).Lte(key, end)
	return f
}

func (f *Find) Contain(field string, key string) *Find {
	f.q.Regex(field, regexp.QuoteMeta(key))
	return f
}

func (f *Find) Regex(field string, pattern string) *Find {
	f.q.Regex(field, pattern)
	return f
}

func (f *Find) Ne(field string, value any) *Find {
	f.q.Ne(field, value)
	return f
}

func (f *Find) Id(id any) *Find {
	f.q.Id(id)
	return f
}

func (f *Find) WithField(fields ...string) *Find {
	f.fields = append(f.fields, fields...)
	return f
}

func (f *Find) WithOptions(opts ...*mongoOptions.FindOptions) *Find {
	f.opts = append(f.opts, opts...)
	return f
}

func (f *Find) WithOneOptions(opts ...*mongoOptions.FindOneOptions) *Find {
	f.oneOpts = append(f.oneOpts, opts...)
	return f
}

func (f *Find) WithReplaceOptions(opts ...*mongoOptions.FindOneAndReplaceOptions) *Find {
	f.replaceOpts = append(f.replaceOpts, opts...)
	return f
}

func (f *Find) WithUpdateOptions(opts ...*mongoOptions.FindOneAndUpdateOptions) *Find {
	f.updateOpts = append(f.updateOpts, opts...)
	return f
}

func (f *Find) WithCountOptions(opts ...*mongoOptions.CountOptions) *Find {
	f.countOpts = append(f.countOpts, opts...)
	return f
}

func (f *Find) Find(ctx context.Context, page, size int64, fieldValues ...any) result {
	// deal with fieldValues
	if len(fieldValues) > 0 {
		if len(fieldValues)%2 != 0 {
			return pairErrResult
		}

		if err := appendPairs(f.q, fieldValues...); err != nil {
			return newErrorResult(err)
		}
	}

	opts := mongoOptions.Find()
	if len(f.sorts) > 0 {
		opts.SetSort(toBsonD(f.sorts))
	}

	if len(f.fields) > 0 {
		opts.SetProjection(toProjection(f.fields...))
	}

	if size > 0 {
		if page > 1 {
			opts.SetSkip((page - 1) * size)
		}
		opts.SetLimit(size)
	}

	if len(f.orConditions) > 0 {
		f.q.Or(f.toInterface(f.orConditions)...)
	}

	f.opts = append(f.opts, opts)

	r, err := f.collection.Find(ctx, f.q.Build(), f.opts...)
	if err != nil {
		return newErrorResult(err)
	}

	err = r.Err()
	if err != nil {
		return newErrorResult(err)
	}

	return newCursorResult(r, ctx)
}

func (f *Find) FindOne(ctx context.Context, fieldValues ...any) result {
	if len(fieldValues) > 0 {
		if len(fieldValues)%2 != 0 {
			return pairErrResult
		}

		if err := appendPairs(f.q, fieldValues...); err != nil {
			return newErrorResult(err)
		}
	}
	r := f.collection.FindOne(ctx, f.q.Build(), f.oneOpts...)
	if r.Err() != nil {
		return newErrorResult(r.Err())
	}

	return newSingleResult(r)
}

func (f *Find) CountDocuments(ctx context.Context, fieldValues ...any) (int64, error) {
	if len(fieldValues) > 0 {
		if err := appendPairs(f.q, fieldValues...); err != nil {
			return 0, err
		}
	}
	return f.collection.CountDocuments(ctx, f.q.Build(), f.countOpts...)
}

func toBsonD(pairs []interface{}) bson.D {
	ret := bson.D{}
	e := bson.E{}
	for i, pair := range pairs {
		if i%2 == 0 {
			e.Key = pair.(string)
		} else {
			e.Value = pair
			ret = append(ret, e)
			e = bson.E{}
		}
	}
	return ret
}

func toProjection(fields ...string) bson.D {
	out := bson.D{}
	idExist := false
	for _, field := range fields {
		if field == "_id" {
			idExist = true
		}
		out = append(out, bson.E{Key: field, Value: 1})
	}
	if !idExist {
		out = append(out, bson.E{Key: "_id", Value: 0})
	}
	return out
}

func (f *Find) toInterface(d []bson.D) []interface{} {
	ret := make([]interface{}, len(d))
	for i, d := range d {
		ret[i] = d
	}
	return ret
}

func appendPairs(q *query.Builder, pairs ...interface{}) error {
	if len(pairs)%2 != 0 {
		return errors.New("pairs must be key-value pairs")
	}

	var key string
	for i, pair := range pairs {
		if i%2 == 0 {
			key = pair.(string)
		} else {
			q.Eq(key, pair)
		}
	}
	return nil
}
