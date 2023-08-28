package entity

import (
	"fmt"
	"testing"
)

func TestAllPred(t *testing.T) {
	fmt.Println(Once(OfCond("one", "=", 1), OfCond("two", "=", 2)).Render())
	fmt.Println(
		And(
			Once(OfCond("one", "=", 1)),
			Once(OfCond("two", "=", 2), OfCond("three", "=", 3)),
		).Render(),
	)
	fmt.Println(
		And(
			Once(OfCond("one", "=", 1)),
			Once(OfCond("two1", "=", 21), OfCond("three1", "=", 31)),
			Or(
				Once(OfCond("one1", "=", 11), OfCond("one2", "=", 12), OfCond("one3", "=", 13)),
				And(Once(OfCond("one4", "=", 14), OfCond("one5", "=", 15), OfCond("one6", "=", 16))),
				Once(OfCond("two12", "=", 212), OfCond("three12", "=", 312)),
			),
			Once(OfCond("two", "=", 2), OfCond("three", "=", 3)),
		).Render(),
	)
}

func BenchmarkAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		TestAllPred(nil)
	}
}
