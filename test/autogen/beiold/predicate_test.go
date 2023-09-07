// Package beiold
// @author tabuyos
// @since 2023/9/5
// @description beiold
package beiold

import (
	"fmt"
	"testing"
)

func TestOnce(t *testing.T) {
	once := Once("a", EQ)

	fmt.Printf("rs:> %+v\n", once)
}

func TestPredicate_And(t *testing.T) {
	pred := Once("a", EQ, 1, 3).And(Once("b", NQ, 2))

	fmt.Printf("rs:> %+v\n", pred)
}

func TestPredicate_Or(t *testing.T) {
	pred := Once("a", EQ, 1, 3).Or(Once("b", NQ, 2))

	fmt.Printf("rs:> %+v\n", pred)
}

func TestPredicate_Mix(t *testing.T) {
	pred := Once("a", EQ, 1, 3).Or(Once("b", NQ, 2)).And(Once("c", EQ, 11, 12))

	fmt.Printf("rs:> %+v\n", pred)
}

func TestPredicate_SQL(t *testing.T) {
	pred := Once("a", EQ, 1, 3).Or(Once("b", NQ, 2)).And(Once("c", EQ, 11, 12))
	sql, values := pred.SQL()

	fmt.Printf("rs:> %+v %+v\n", sql, values)
}

func TestPredicate_PureSQL(t *testing.T) {
	pred := Once("a", EQ, 1, 3).Or(Once("b", NQ, 2)).And(Once("c", EQ, 11, 12), Once("d", EQ, 21, 22))
	sql, values := pred.SQL()

	fmt.Printf("rs:> %+v %+v\n", sql, values)
}

func TestPredicate_PureSQL1(t *testing.T) {
	pred := Once("a", EQ, 1, 3).Or(Once("b", NQ, 2), Once("bb", NQ, 2)).And(Once("c", EQ, 11, 12), Once("d", EQ, 21, 22))
	sql, values := pred.SQL()

	fmt.Printf("rs:> %+v %+v\n", sql, values)
}

func TestPredicate_PureSQL2(t *testing.T) {
	pred := Once("a", EQ, 1, 3).Or(Once("b", NQ, 2), Once("bb", NQ, 2))
	sql, values := pred.SQL()

	fmt.Printf("rs:> %+v %+v\n", sql, values)
}

func TestPredicate_PureSQL3(t *testing.T) {
	pred := Once("a", EQ, 1, 3).Or(Once("b", NQ, 2), Once("b", NQ, 2)).And(Once("c", EQ, 11, 12), Once("c", EQ, 11, 12))
	sql, values := pred.SQL()

	fmt.Printf("rs:> %+v %+v\n", sql, values)
	fmt.Printf("rs:> %+v %+v\n", pred.String(), values)
}

func TestPredicate_PureSQL4(t *testing.T) {
	pred := Once("a", EQ, 1).Or(Once("b", NQ, "fdsa"), Once("b", NQ, 2)).And(Once("c", EQ, 11, 12), Once("c", EQ, 11, 12))

	fmt.Printf("rs:> %+v\n", pred.String())
}
