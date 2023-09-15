// Package redisx
// @author tabuyos
// @since 2023/8/7
// @description redisx
package redisx

import (
	"deepsea/config"
	"deepsea/config/env"
	"deepsea/helper/recorderx"
	"fmt"
	"testing"
	"time"
)

func TestRedisXLocker_Lock(t *testing.T) {
	env.SpecialEnv("dev")
	config.InitConfig()
	InitRedisX()

	recorder := recorderx.DefaultRecorder()
	locker := NewRedisLocker(recorder)
	locker.Lock("test-1")
	defer locker.Unlock()

	fmt.Println("b...")
	time.Sleep(2 * time.Second)
	fmt.Println("a...")
}
