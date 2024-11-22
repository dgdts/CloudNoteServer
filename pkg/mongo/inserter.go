package mongo

import (
	"context"
	"errors"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Insert struct {
	collection *mongoDriver.Collection
	opts       []*mongoOptions.InsertManyOptions
	oneOpts    []*mongoOptions.InsertOneOptions
}

func Inserter(collection string, opts ...CollectionOption) *Insert {
	return &Insert{
		collection: RawCollection(collection, opts...),
	}
}

func (i *Insert) WithInsertManyOptions(opts ...*mongoOptions.InsertManyOptions) {
	i.opts = opts
}

func (i *Insert) WithInsertOneOptions(opts ...*mongoOptions.InsertOneOptions) {
	i.oneOpts = opts
}

func (i *Insert) Insert(ctx context.Context, documents ...interface{}) result {
	l := len(documents)
	if l == 0 {
		return &errorResult{
			err: errors.New("documents is empty"),
		}
	}
	if l == 1 {
		return i.insertOne(ctx, documents[0])
	}
	return i.insertMany(ctx, documents...)
}

func (i *Insert) insertOne(ctx context.Context, document interface{}) result {
	r, err := i.collection.InsertOne(ctx, document, i.oneOpts...)
	if err != nil {
		return newErrorResult(err)
	}
	return &insertOneResult{
		result: r,
	}
}

func (i *Insert) insertMany(ctx context.Context, documents ...interface{}) result {
	r, err := i.collection.InsertMany(ctx, documents, i.opts...)
	if err != nil {
		return newErrorResult(err)
	}
	return &insertManyResult{
		result: r,
	}
}
