// Package middleware
// @author tabuyos
// @since 2023/8/6
// @description middleware
package middleware

import (
	"deepsea/helper/security"
	"github.com/gin-gonic/gin"
	"slices"
)

type RbacHandler Handler

type rbac struct {
	keys []string
	// mode 模式, and: 0, or: 1 default and
	mode int8
}

const (
	MAnd  = 0
	MOr   = 1
	TRole = 0
	TPerm = 1
)

type role rbac

type perm rbac

type and struct {
	chain []RbacHandler
}

type or struct {
	chain []RbacHandler
}

func NewRole(r string) RbacHandler {
	return &role{[]string{r}, MAnd}
}

func NewAnyRole(rs ...string) RbacHandler {
	return &role{rs, MOr}
}

func NewAllRole(rs ...string) RbacHandler {
	return &role{rs, MAnd}
}

func NewPerm(r string) RbacHandler {
	return &perm{[]string{r}, MAnd}
}

func NewAnyPerm(rs ...string) RbacHandler {
	return &perm{rs, MOr}
}

func NewAllPerm(rs ...string) RbacHandler {
	return &perm{rs, MAnd}
}

func handleRbac(ctx *gin.Context, ty int8, keys []string, mode int8) bool {
	if len(keys) == 0 {
		return true
	}
	var rps []string
	if ty == TRole {
		rps = security.GetRoleFromCtx(ctx)
	}
	if ty == TPerm {
		rps = security.GetPermFromCtx(ctx)
	}
	// and
	if mode == 0 {
		for _, key := range keys {
			ok := slices.Contains(rps, key)
			if !ok {
				return false
			}
		}
		return true
	} else {
		for _, key := range keys {
			ok := slices.Contains(rps, key)
			if ok {
				return true
			}
		}
		return false
	}
}

func (rec *role) Handle(ctx *gin.Context) bool {
	return handleRbac(ctx, TRole, rec.keys, rec.mode)
}

func (rec *perm) Handle(ctx *gin.Context) bool {
	return handleRbac(ctx, TPerm, rec.keys, rec.mode)
}

func (rec *and) Handle(ctx *gin.Context) bool {
	if len(rec.chain) == 0 {
		return true
	}
	for _, handler := range rec.chain {
		hs := handler.Handle(ctx)
		if !hs {
			return false
		}
	}
	return true
}

func (rec *or) Handle(ctx *gin.Context) bool {
	if len(rec.chain) == 0 {
		return true
	}
	for _, handler := range rec.chain {
		hs := handler.Handle(ctx)
		if hs {
			return true
		}
	}
	return false
}

// And 全称, 必须所有的 handler 全部满足要求
func And(chain ...RbacHandler) RbacHandler {
	return &and{chain}
}

// Or 存在, 只要存在一个 handler 满足要求即可
func Or(chain ...RbacHandler) RbacHandler {
	return &or{chain}
}
