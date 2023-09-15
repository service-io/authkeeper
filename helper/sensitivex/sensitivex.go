// Package sensitivex
// @author tabuyos
// @since 2023/8/15
// @description sensitivex
package sensitivex

import "deepsea/model/entity"

func EraseSensitive[T ~[]*E, E any](ls T, fn func(*E)) {
	for _, e := range ls {
		fn(e)
	}
}

func EraseAccountsSensitive(accounts ...*entity.PlatAccount) {
	EraseSensitive(accounts, func(account *entity.PlatAccount) {
		account.Pwd = nil
	})
}
