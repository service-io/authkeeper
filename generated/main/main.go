// Package main
// @author tabuyos
// @since 2023/8/29
// @description main
package main

import (
	"metis/generated/baseentity"
	"metis/generated/entity"
)

func main() {
	entity.New(nil).RenderSelf()
	baseentity.New(nil).RenderAuto()
}
