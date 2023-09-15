// Package redisx
// @author tabuyos
// @since 2023/8/5
// @description redisx
package redisx

import (
	"context"
	"deepsea/config/constant"
	"deepsea/helper"
	"deepsea/helper/recorderx"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
	"time"
)

// redisXEmitterPool redis emitter 池
var redisXEmitterPool = &sync.Pool{New: func() interface{} {
	return new(redisXEmitter)
}}

type IRedisXEmitter interface {
	BuildKey(seg ...string) string
	BuildVal(seg ...string) string

	Client() redis.UniversalClient

	Pipeliner() redis.Pipeliner
	TxPipeliner() redis.Pipeliner

	Watch(fn func(*redis.Tx) error, keys ...string) error

	Expire(key string, expiration time.Duration) bool

	Del(keys ...string) bool

	Exists(keys ...string) bool

	Get(key string) string
	Set(key string, value interface{}, expiration time.Duration) bool
	SetNX(key string, value interface{}, expiration time.Duration) bool
	SetXX(key string, value interface{}, expiration time.Duration) bool

	SAdd(key string, value ...interface{}) bool
	SMembers(key string) []string
}

type redisXEmitter struct {
	client   redis.UniversalClient
	timeout  time.Duration
	recorder func(err error)
}

func NewRedisEmitter(recorder recorderx.Recorder) (IRedisXEmitter, func()) {
	return NewRedisEmitterWithTimeout(3*time.Second, recorder)
}

func NewRedisEmitterWithTimeout(timeout time.Duration, recorder recorderx.Recorder) (IRedisXEmitter, func()) {
	var errorHandler = func(err error) {
		if err != nil {
			panic(err)
		}
	}
	if recorder != nil {
		logger := &redisLogger{recorder}
		redis.SetLogger(logger)
		errorHandler = helper.ErrToLog(recorder)
	}

	emitter := redisXEmitterPool.Get().(*redisXEmitter)
	emitter.client = FetchRedisX()
	emitter.timeout = timeout
	emitter.recorder = errorHandler
	rel := func() {
		if emitter.client == nil {
			return
		}
		emitter.client = nil
		emitter.recorder = nil
		emitter.timeout = 0
		redisXEmitterPool.Put(emitter)
	}
	return emitter, rel
}

func handleInvalidate() {

}

func (rec *redisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	rec.recorder.WithOptions(recorderx.AddCallerSkip(2)).Info(fmt.Sprintf(format, v...))
}

func (rec *redisXEmitter) getCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), rec.timeout)
}

func (rec *redisXEmitter) BuildKey(seg ...string) string {
	return BuildKey(seg...)
}

func (rec *redisXEmitter) BuildVal(seg ...string) string {
	return BuildVal(seg...)
}

func BuildKey(seg ...string) string {
	return strings.Join(seg, constant.KeyDelimiter)
}

func BuildVal(seg ...string) string {
	return strings.Join(seg, constant.ValDelimiter)
}

func (rec *redisXEmitter) Client() redis.UniversalClient {
	return rec.client
}

func (rec *redisXEmitter) Pipeliner() redis.Pipeliner {
	return rec.Client().Pipeline()
}

func (rec *redisXEmitter) TxPipeliner() redis.Pipeliner {
	return rec.Client().TxPipeline()
}

func (rec *redisXEmitter) Watch(fn func(*redis.Tx) error, keys ...string) error {
	ctx, cancel := rec.getCtx()
	defer cancel()
	return rec.Client().Watch(ctx, fn, keys...)
}

func (rec *redisXEmitter) Expire(key string, expiration time.Duration) bool {
	ctx, cancel := rec.getCtx()
	defer cancel()

	result, err := rec.Client().Expire(ctx, key, expiration).Result()
	rec.recorder(err)
	return result
}

func (rec *redisXEmitter) Del(keys ...string) bool {
	ctx, cancel := rec.getCtx()
	defer cancel()

	_, err := rec.Client().Del(ctx, keys...).Result()
	rec.recorder(err)
	return true
}

func (rec *redisXEmitter) Exists(keys ...string) bool {
	ctx, cancel := rec.getCtx()
	defer cancel()

	result, err := rec.Client().Exists(ctx, keys...).Result()
	rec.recorder(err)
	return result == int64(len(keys))
}

func (rec *redisXEmitter) Get(key string) string {
	ctx, cancel := rec.getCtx()
	defer cancel()

	result, err := rec.Client().Get(ctx, key).Result()
	rec.recorder(err)
	return result
}

func (rec *redisXEmitter) Set(key string, value interface{}, expiration time.Duration) bool {
	ctx, cancel := rec.getCtx()
	defer cancel()

	result, err := rec.Client().Set(ctx, key, value, expiration).Result()
	rec.recorder(err)

	return result == "OK"
}

func (rec *redisXEmitter) SetNX(key string, value interface{}, expiration time.Duration) bool {
	ctx, cancel := rec.getCtx()
	defer cancel()

	result, err := rec.Client().SetNX(ctx, key, value, expiration).Result()
	rec.recorder(err)

	return result
}

func (rec *redisXEmitter) SetXX(key string, value interface{}, expiration time.Duration) bool {
	ctx, cancel := rec.getCtx()
	defer cancel()

	result, err := rec.Client().SetXX(ctx, key, value, expiration).Result()
	rec.recorder(err)

	return result
}

func (rec *redisXEmitter) SAdd(key string, value ...interface{}) bool {
	ctx, cancel := rec.getCtx()
	defer cancel()

	counter, err := rec.Client().SAdd(ctx, key, value...).Result()
	rec.recorder(err)

	return int64(len(value)) == counter
}

func (rec *redisXEmitter) SMembers(key string) []string {
	ctx, cancel := rec.getCtx()
	defer cancel()

	rs, err := rec.Client().SMembers(ctx, key).Result()
	rec.recorder(err)

	return rs
}
