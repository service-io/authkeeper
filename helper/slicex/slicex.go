// Package slicex
// @author tabuyos
// @since 2023/8/14
// @description slicex
package slicex

func Grouping[T ~[]E, K comparable, E any](ls T, keyFn func(E) K) map[K]T {
	rs := make(map[K]T)
	for _, e := range ls {
		c := rs[keyFn(e)]
		if c == nil {
			rs[keyFn(e)] = T{e}
		} else {
			rs[keyFn(e)] = append(c, e)
		}
	}
	return rs
}

func ToTree[T ~[]E, K comparable, E any](ls T, idFn func(E) K, pidFn func(E) K, childrenFn func(E, T)) map[K]T {
	rs := make(map[K]T)
	grouping := Grouping(ls, pidFn)

	for _, e := range ls {
		childrenFn(e, grouping[idFn(e)])
		rs[pidFn(e)] = append(rs[pidFn(e)], e)
	}

	return rs
}

func ToFlat() {

}
