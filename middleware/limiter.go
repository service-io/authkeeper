// Package middleware
// @author tabuyos
// @since 2023/8/14
// @description middleware
package middleware

import (
	"deepsea/helper/security"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map/v2"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type ApiLimiter interface {
	Get(ctx *gin.Context) *rate.Limiter
}

type tl struct {
	limiter *rate.Limiter
	timer   *time.Timer
}

type sfn func() (*rate.Limiter, *time.Timer)

type apiLimiter struct {
	container cmap.ConcurrentMap[string, *tl]
	supplier  sfn
	timer     *time.Timer
	mu        sync.Mutex
}

type userModifyApiLimiter apiLimiter

func NewUserModifyApiLimiter(supplier sfn) ApiLimiter {
	return &userModifyApiLimiter{
		container: cmap.New[*tl](),
		supplier:  supplier,
		mu:        sync.Mutex{},
	}
}

// NewDefaultUserModifyApiLimiter 修改接口, 一秒一次, 十秒过期, 自动续签
func NewDefaultUserModifyApiLimiter() ApiLimiter {
	supplier := func() (*rate.Limiter, *time.Timer) {
		return rate.NewLimiter(rate.Every(1*time.Second), 1), time.NewTimer(10 * time.Second)
	}
	return NewUserModifyApiLimiter(supplier)
}

func (rec *userModifyApiLimiter) Get(ctx *gin.Context) *rate.Limiter {
	rec.mu.Lock()
	defer rec.mu.Unlock()
	// recorder := recorderx.FetchRecorder(ctx)
	id := security.GetAccountIDString(ctx)
	tls, ok := rec.container.Get(id)
	if ok {
		_ = tls.timer.Reset(10 * time.Second)
		return tls.limiter
	}

	limiter, timer := rec.supplier()
	rec.container.Set(id, &tl{limiter, timer})

	go func() {
		select {
		case <-timer.C:
			// recorder.Info(fmt.Sprintf("移除前容器大小: %d, 当前待移除的 key: %s\n", len(rec.container.Items()), id))
			rec.container.Remove(id)
			// recorder.Info(fmt.Sprintf("移除后容器大小: %d\n", len(rec.container.Items())))
		}
	}()

	return limiter
}
