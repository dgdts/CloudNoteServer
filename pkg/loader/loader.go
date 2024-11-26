package loader

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/sync/singleflight"
)

type Func[T any] func(ctx context.Context, key string) (T, error)

type ChainFunc[T any] []Func[T]

var ErrNext = errors.New("next error")
var ErrNoResult = errors.New("no result")

func (c ChainFunc[T]) Load(ctx context.Context, key string) (T, error) {
	var zero T

	for _, f := range c {
		result, err := f(ctx, key)
		if err != nil {
			if errors.Is(err, ErrNext) {
				continue
			}
			return zero, err
		}
		return result, nil
	}
	return zero, ErrNoResult
}

var g = singleflight.Group{}
var keyMap = map[string]struct{}{}

func SingleFlightLoader[T any](key string, f Func[T]) Func[T] {
	_, ok := keyMap[key]
	if ok {
		panic("key already exists")
	}
	keyMap[key] = struct{}{}

	return func(ctx context.Context, key string) (T, error) {
		r, err, _ := g.Do(key, func() (interface{}, error) {
			s := time.Now()
			defer func() {
				cost := time.Since(s)
				hlog.Infof("SingleFlightLoader: %s, cost: %v", key, cost)
			}()
			return f(ctx, key)
		})
		if err != nil {
			var zero T
			return zero, err
		}
		return r.(T), nil
	}
}
