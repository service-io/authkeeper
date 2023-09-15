// Package redisx
// @author tabuyos
// @since 2023/8/7
// @description redisx
package redisx

import (
	"context"
	"deepsea/config/constant"
	"deepsea/helper"
	"deepsea/helper/concurrency"
	"deepsea/helper/recorderx"
	"deepsea/helper/runerror"
	"errors"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"time"
)

type IRedisXLocker interface {
	Lock(key string)
	Unlock()
}

type redisXLocker struct {
	key      string
	ctx      context.Context
	lock     *redislock.Lock
	recorder func(err error)
}

func NewRedisLocker(recorder recorderx.Recorder) IRedisXLocker {
	if recorder != nil {
		logger := &redisLogger{recorder}
		redis.SetLogger(logger)
	}
	return &redisXLocker{recorder: helper.ErrToLogAndPanic(recorder)}
}

func (rec *redisXLocker) Lock(key string) {
	if len(key) == 0 {
		panic(runerror.NewAll(runerror.GetSysErp(runerror.ModCache, runerror.LenError), errors.New("lock key to short")))
	}

	mutexLock := concurrency.NewMutexLockRoutine()
	mutexLock.Lock()
	defer mutexLock.Unlock()

	ctx := context.Background()
	rec.key = key
	rec.ctx = ctx

	locker := redislock.New(FetchRedisX())
	backoff := redislock.LimitRetry(redislock.LinearBackoff(100*time.Millisecond), 50)

	lock, err := locker.Obtain(ctx, BuildKey(constant.DistributedLockPrefix, key), 10*time.Second, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err == redislock.ErrNotObtained {
		rec.recorder(err)
	} else if err != nil {
		rec.recorder(err)
	}

	rec.lock = lock
}

func (rec *redisXLocker) Unlock() {
	err := rec.lock.Release(rec.ctx)
	rec.recorder(err)
}
