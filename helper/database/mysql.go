// Package database
// @author tabuyos
// @since 2023/6/30
// @description database
package database

import (
	"context"
	"database/sql"
	"deepsea/config"
	"deepsea/config/constant"
	"deepsea/helper/recorderx"
	"github.com/go-sql-driver/mysql"
	"github.com/qustavo/sqlhooks/v2"
	"log/slog"
	"strings"
	"time"
)

var db *sql.DB

type Hooks struct{}

// Before hook will print the query with it's args and return the context with the timestamp
func (h *Hooks) Before(ctx context.Context, _ string, _ ...interface{}) (context.Context, error) {
	return context.WithValue(ctx, "begin", time.Now()), nil
}

// After hook will get the timestamp registered on the Before hook and print the elapsed time
func (h *Hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	watchSQL(ctx, query, args...)
	return ctx, nil
}

func (h *Hooks) OnError(ctx context.Context, err error, query string, args ...interface{}) error {
	watchSQL(ctx, query, args...)
	return err
}

func watchSQL(ctx context.Context, query string, args ...interface{}) {
	if strings.Contains(query, "sys_log") {
		return
	}
	begin := ctx.Value("begin").(time.Time)
	recorder := recorderx.DefaultRecorder()
	traceID := ctx.Value(constant.TraceIdKey)
	tid, ok := traceID.(string)
	if !ok {
		tid = ""
	}
	if traceID != nil && tid != "" {
		recorder = recorder.With(slog.String(constant.TraceIdKey, tid))
	}
	recorder.Infof("sql: %s, bvs: %#v, took: %s", query, args, time.Since(begin))
}

func InitDB() {
	recorder := recorderx.DefaultRecorder()

	recorder.Info("初始化 MySQL...")
	recorder.Info("添加 MySQL 驱动...")

	var driverName = "mysqlWithHooks"

	sql.Register(driverName, sqlhooks.Wrap(&mysql.MySQLDriver{}, &Hooks{}))

	tomlConfig := config.TomlConfig()
	single := tomlConfig.MySQL.Single

	var params = make(map[string]string)
	params["parseTime"] = "true"
	params["loc"] = "Asia/Shanghai"

	cfg := mysql.Config{
		User:                 single.Username,
		Passwd:               single.Password,
		Net:                  "tcp",
		Addr:                 single.Addr,
		DBName:               single.DbName,
		AllowNativePasswords: true,
		Params:               params,
	}

	var err error
	db, err = sql.Open(driverName, cfg.FormatDSN())
	if err != nil {
		recorder.Error(err.Error())
	}
	err = db.Ping()
	if err != nil {
		recorder.Info("MySQL 初始化失败...")
		panic(err)
	}
	recorder.Info("MySQL 初始化完成...")
}

func FetchDB() *sql.DB {
	return db
}
