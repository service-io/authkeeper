// Package tree
// @author tabuyos
// @since 2023/8/15
// @description tree
package tree

import "deepsea/model/dto"

type INode[T ~[]E, K comparable, E any] interface {
	ID() K
	Pid() K
	IsRoot() bool
	Children(T)
}

type ResourceNode struct {
	*dto.PlatResource
	Children []*ResourceNode `json:"children"`
}

type CommonAreaNode struct {
	*dto.CommonArea
	Children []*CommonAreaNode `json:"children"`
}

type RoleNode struct {
	*dto.PlatRole
	Children []*RoleNode `json:"children"`
}

type TenantDeptNode struct {
	*dto.TenantDept
	Children []*TenantDeptNode `json:"children"`
}

type TenantPostNode struct {
	*dto.TenantPost
	Children []*TenantPostNode `json:"children"`
}
