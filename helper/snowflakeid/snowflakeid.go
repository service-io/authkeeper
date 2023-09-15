// Package snowflakeid
// @author tabuyos
// @since 2023/8/7
// @description snowflake
package snowflakeid

import (
	"deepsea/config"
	"deepsea/helper/concurrency"
	"deepsea/helper/recorderx"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
var lock concurrency.MutexLock

func InitSnowflake() {
	recorder := recorderx.DefaultRecorder()
	recorder.Info("初始化雪花算法...")
	snowflakeConfig := config.TomlConfig().Snowflake
	nodeId := snowflakeConfig.Node
	newNode, err := snowflake.NewNode(int64(nodeId))
	if err != nil {
		recorder.Info("雪花算法初始化失败...")
		panic(err)
	}
	node = newNode
	lock = concurrency.NewMutexLockRoutine()
	recorder.Info("雪花算法初始化成功...")
}

func Generate() int64 {
	lock.Lock()
	defer lock.Unlock()

	id := node.Generate()
	return id.Int64()
}

func GeneratePtr() *int64 {
	id := Generate()
	return &id
}
