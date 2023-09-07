// Package beiold
// @author tabuyos
// @since 2023/9/6
// @description beiold
package beiold

import (
	"fmt"
	"testing"
)

func TestRefTable_SQL(t *testing.T) {
	table := OfRef("tabuyos").As("ts")
	account := OfRef("account").As("at")
	user := OfRef("user").As("ur")
	role := OfRef("role").As("rl")

	sql := table.Ref(
		account.JoinType(LeftJoin).On(table.RefKey("name"), EQ, account.RefKey("nickname")),
		user.JoinType(LeftJoin).On(account.RefKey("age"), EQ, user.RefKey("year")).Ref(
			role.JoinType(LeftJoin).On(user.RefKey("role_id"), EQ, role.RefKey("id")),
		),
	).SQL()

	fmt.Println(len(table.FlatAll()))
	fmt.Printf("rs:> %+v\n", sql)
}
