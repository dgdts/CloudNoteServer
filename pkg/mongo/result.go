package mongo

import (
	"context"
	"encoding/json"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type result interface {
	Read(obj interface{}) error
	Error() error
}

// insertOneResult
var _ result = (*insertOneResult)(nil)

type insertOneResult struct {
	result *mongo.InsertOneResult
}

func (ir *insertOneResult) Error() error {
	return nil
}

func (ir *insertOneResult) Read(obj interface{}) error {
	data, err := json.Marshal(ir.result.InsertedID)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}

// insertManyResult
var _ result = (*insertManyResult)(nil)

type insertManyResult struct {
	result *mongo.InsertManyResult
}

func (ir *insertManyResult) Error() error {
	return nil
}

func (ir *insertManyResult) Read(obj interface{}) error {
	data, err := json.Marshal(ir.result.InsertedIDs)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}

// cursorResult
var _ result = (*cursorResult)(nil)

type cursorResult struct {
	cursor *mongo.Cursor
	ctx    context.Context
}

func (cr *cursorResult) Error() error {
	return cr.cursor.Err()
}

func (cr *cursorResult) Read(obj interface{}) error {
	if cr.cursor.Err() != nil {
		return cr.cursor.Err()
	}
	defer cr.cursor.Close(cr.ctx)
	return cr.cursor.All(cr.ctx, obj)
}

func newCursorResult(cursor *mongo.Cursor, ctx context.Context) result {
	return &cursorResult{
		cursor: cursor,
		ctx:    ctx,
	}
}

// singleResult
var _ result = (*singleResult)(nil)

type singleResult struct {
	result *mongo.SingleResult
}

func (sr *singleResult) Error() error {
	return sr.result.Err()
}

func (sr *singleResult) Read(obj interface{}) error {
	if sr.result.Err() != nil {
		return sr.result.Err()
	}
	return sr.result.Decode(obj)
}

func newSingleResult(result *mongo.SingleResult) result {
	return &singleResult{
		result: result,
	}
}

// updateResultInterface
type updateResultInterface interface {
	result
	Affected() int64
}

// updateResult
var _ updateResultInterface = (*updateResult)(nil)

type updateResult struct {
	result *mongo.UpdateResult
}

func (ur *updateResult) Error() error {
	return nil
}

func (ur *updateResult) Read(obj interface{}) error {
	d, err := json.Marshal(ur.result.ModifiedCount)
	if err != nil {
		return err
	}
	return json.Unmarshal(d, ur.result)
}

func newUpdateResult(result *mongo.UpdateResult) updateResultInterface {
	return &updateResult{
		result: result,
	}
}

func (ur *updateResult) Affected() int64 {
	return ur.result.ModifiedCount + ur.result.UpsertedCount
}

// errorResult
var _ updateResultInterface = (*errorResult)(nil)

type errorResult struct {
	err error
}

func (er *errorResult) Error() error {
	return er.err
}

func (er *errorResult) Read(obj interface{}) error {
	return er.err
}

func (er *errorResult) Affected() int64 {
	return 0
}

func newErrorResult(err error) updateResultInterface {
	return &errorResult{
		err: err,
	}
}

var ErrParamMustPairs = errors.New("fieldValues must be even")
var pairErrResult = newErrorResult(ErrParamMustPairs)
