// Package config
// @author tabuyos
// @since 2023/6/30
// @description config
package config

import (
	"bytes"
	"deepsea/config/env"
	"deepsea/helper/file"
	_ "embed"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"time"
)

var config = new(Config)

type MySQL struct {
	Read   DB   `toml:"read"`
	Write  DB   `toml:"write"`
	Single DB   `toml:"single"`
	Base   Base `toml:"base"`
}

type Base struct {
	MaxOpenConn     int           `toml:"maxOpenConn"`
	MaxIdleConn     int           `toml:"maxIdleConn"`
	ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
}

type DB struct {
	Addr     string `toml:"addr"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DbName   string `toml:"dbName"`
}

type Redis struct {
	Addrs        []string `toml:"addrs"`
	Username     string   `toml:"username"`
	Password     string   `toml:"password"`
	Db           int      `toml:"db"`
	MaxRetries   int      `toml:"maxRetries"`
	PoolSize     int      `toml:"poolSize"`
	MinIdleConns int      `toml:"minIdleConns"`
}

type Mail struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	To   string `toml:"to"`
}

type DingTalk struct {
}

type Security struct {
	Secret     string        `toml:"secret"`
	Length     int           `toml:"length"`
	AccessTTL  time.Duration `toml:"accessTtl"`
	RefreshTTL time.Duration `toml:"refreshTtl"`
}

type Language struct {
	Local string `toml:"local"`
}

type Code struct {
	App    string   `toml:"app"`
	Module []string `toml:"module"`
}

type Server struct {
	Address string `toml:"address"`
	IP      string `toml:"ip"`
	Port    string `toml:"port"`
}

type Snowflake struct {
	Node int `toml:"node"`
}

type OSS struct {
	Aliyun AliyunOSS `toml:"aliyun"`
}

type AliyunOSS struct {
	Endpoint        string `toml:"endpoint"`
	AccessKey       string `toml:"accessKey"`
	Secret          string `toml:"secret"`
	BucketName      string `toml:"bucketName"`
	SignTTL         int    `toml:"signTtl"`
	Prefix          string `toml:"prefix"`
	Region          string `toml:"region"`
	Domain          string `toml:"domain"`
	Scheme          string `toml:"scheme"`
	RoleArn         string `toml:"roleArn"`
	RoleSessionName string `toml:"roleSessionName"`
}

type Config struct {
	MySQL     MySQL     `toml:"mysql"`
	Redis     Redis     `toml:"redis"`
	Mail      Mail      `toml:"mail"`
	DingTalk  DingTalk  `toml:"dingtalk"`
	Security  Security  `toml:"security"`
	Language  Language  `toml:"language"`
	Server    Server    `toml:"server"`
	Snowflake Snowflake `toml:"snowflake"`
	OSS       OSS       `toml:"oss"`
}

var (
	//go:embed dev_configs.toml
	devConfigs []byte

	//go:embed fat_configs.toml
	fatConfigs []byte

	//go:embed uat_configs.toml
	uatConfigs []byte

	//go:embed pro_configs.toml
	proConfigs []byte
)

func InitConfig() {
	var reader io.Reader

	switch env.Active().Value() {
	case "dev":
		reader = bytes.NewReader(devConfigs)
	case "fat":
		reader = bytes.NewReader(fatConfigs)
	case "uat":
		reader = bytes.NewReader(uatConfigs)
	case "pro":
		reader = bytes.NewReader(proConfigs)
	default:
		reader = bytes.NewReader(fatConfigs)
	}

	viper.SetConfigType("toml")

	if err := viper.ReadConfig(reader); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.AddConfigPath("./config")

	configFile := "./config/" + env.Active().Value() + "_configs.toml"

	_, ok := file.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
			}
		}(f)

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func TomlConfig() Config {
	return *config
}
