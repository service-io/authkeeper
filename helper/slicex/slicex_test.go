// Package slicex
// @author tabuyos
// @since 2023/8/14
// @description slicex
package slicex

import (
	"deepsea/helper/encodingx"
	"fmt"
	"testing"
)

type child struct {
	pid  int
	name string
}

func TestAppend(t *testing.T) {
	var v []string = nil

	_ = append(v, "ta")

	fmt.Printf("v: %#v\n", v)
}

func TestGrouping(t *testing.T) {
	childs := []child{
		{
			pid:  0,
			name: "c0",
		},
		{
			pid:  0,
			name: "c1",
		},
		{
			pid:  0,
			name: "c2",
		},
		{
			pid:  1,
			name: "c3",
		},
		{
			pid:  1,
			name: "c4",
		},
		{
			pid:  2,
			name: "c5",
		},
	}

	rs := Grouping(childs, func(e child) int {
		return e.pid
	})

	fmt.Printf("分组结果: %#v\n", rs)
	fmt.Printf("分组长度: %#v\n", len(rs))
	fmt.Printf("--------------------\n")
	for k, c := range rs {
		fmt.Printf("分组KEY: %#v\n", k)
		fmt.Printf("分组VAL: %#v\n", c)
	}
}

type Node struct {
	Id       int     `json:"id"`
	ParentId int     `json:"parent_id"`
	Name     string  `json:"name"`
	Children []*Node `json:"children"`
}

func TestToTree(t *testing.T) {
	list := []*Node{
		{1, 0, "A", nil},
		{2, 1, "AA", nil},
		{3, 1, "AB", nil},
		{4, 3, "ABA", nil},
		{5, 3, "ABB", nil},
	}

	tree := ToTree(list, func(e *Node) int {
		return e.Id
	}, func(e *Node) int {
		return e.ParentId
	}, func(e *Node, ns []*Node) {
		e.Children = ns
	})
	nodes := tree[0]
	fmt.Printf("nodes: %+v\n", nodes)
	fmt.Printf("len: %+v\n", len(nodes))
	for _, node := range nodes {
		fmt.Printf("node: %+v\n", node)
	}
	encodingx.InitEncodingX()
	fmt.Printf("json: %#v\n", encodingx.ToJSON(nodes))
}

// List 结构体
type List struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Pid      int    `json:"pid"`
	Children []List `json:"children"`
}

// 数据
var data = []List{
	{Name: "李四", Id: 2, Pid: 0},  //  []
	{Name: "王五", Id: 3, Pid: 0},  // []
	{Name: "赵六", Id: 4, Pid: 3},  // []
	{Name: "吗六", Id: 9, Pid: 3},  // []
	{Name: "张三", Id: 7, Pid: 9},  // []
	{Name: "张五", Id: 10, Pid: 4}, // []
}

/**
 * 递归模式，数组转tree
 * @param arr 目标数组
 * @param pid 第一级 目标id
 * @returns {*[]} tree
 * @constructor
 */
func ArrayToTree(arr []List, pid int) []List {
	var newArr []List
	for _, v := range arr {
		if v.Pid == pid {
			v.Children = ArrayToTree(arr, v.Id)
			newArr = append(newArr, v)
		}
	}
	return newArr
}
