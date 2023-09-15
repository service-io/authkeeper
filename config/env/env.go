// Package env
// @author tabuyos
// @since 2023/7/7
// @description config
package env

import (
	"flag"
	"fmt"
	"strings"
)

var (
	// 激活的环境
	active Environment
	// DEV(Development) 开发环境
	dev Environment = &environment{value: "dev"}
	// FAT(Feature Acceptance Test) 测试环境
	fat Environment = &environment{value: "fat"}
	// UAT(Username Acceptance Test) 预上线环境
	uat Environment = &environment{value: "uat"}
	// PRO(Production) 生产环境
	pro Environment = &environment{value: "pro"}
)

type environment struct {
	value string
}

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsFat() bool
	IsUat() bool
	IsPro() bool
}

func (e *environment) Value() string {
	return e.value
}

func (e *environment) IsDev() bool {
	return e.value == "dev"
}

func (e *environment) IsFat() bool {
	return e.value == "fat"
}

func (e *environment) IsUat() bool {
	return e.value == "uat"
}

func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func SpecialEnv(e string) {
	switch strings.ToLower(strings.TrimSpace(e)) {
	case "dev":
		active = dev
	case "fat":
		active = fat
	case "uat":
		active = uat
	case "pro":
		active = pro
	default:
		active = fat
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'fat' will be used.")
	}
}

func InitEnv() {
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n pro:正式环境\n")
	flag.Parse()

	SpecialEnv(*env)
}

// Active 当前配置的 env
func Active() Environment {
	return active
}
