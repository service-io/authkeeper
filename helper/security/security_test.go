// Package security
// @author tabuyos
// @since 2023/8/8
// @description security
package security

import (
	"deepsea/config"
	"deepsea/config/env"
	"fmt"
	"testing"
	"time"
)

func TestGeneratePassword(t *testing.T) {
	password := GeneratePassword("123456")
	fmt.Println(password)

	ok := ComparePassword(password, "123456s")
	fmt.Println(ok)
}

func TestTime(t *testing.T) {
	env.SpecialEnv("dev")
	config.InitConfig()
	securityConfig := config.TomlConfig().Security
	println(securityConfig.AccessTTL.String())

	duration, err := time.ParseDuration("35h")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(duration.String())
	fmt.Println(duration / time.Hour)
}

func TestNilAddr(t *testing.T) {
	var names []string = nil
	fmt.Println(names)
	fmt.Println(&names)
}

func TestGenTokenInfo(t *testing.T) {
	hash := GenerateHash("123456")
	info := GenTokenInfo("张三", hash)

	at := info.AccessToken

	fmt.Println(at)

	at = nil

	fmt.Println(*info.AccessToken)
	fmt.Println(len(*info.AccessToken))
}
