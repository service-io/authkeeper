// Package sonar
// @author tabuyos
// @since 2023/8/5
// @description sonar 缓存扫描, 获取, 兜底
package sonar

import "errors"

type Sonar[K, V any] interface {
	Backing(rfn resume[K, V], wfn writer[K, V]) Sonar[K, V]
	MissResume(rfn resume[K, V]) Sonar[K, V]
	WriteWith(wfn writer[K, V]) Sonar[K, V]
	Get() *V
	OrDft(dv *V) (*V, error)
}

type resume[K, V any] func(key *K) *V
type writer[K, V any] func(key *K, val *V)

type pair[K, V any] struct {
	resume[K, V]
	writer[K, V]
}

type sonar[K, V any] struct {
	key   *K
	val   *V
	chain []*pair[K, V]
}

func NewSonar[K, V any]() Sonar[K, V] {
	return &sonar[K, V]{}
}

func Lookup[K, V any](fn resume[K, V], key *K) Sonar[K, V] {
	rec := &sonar[K, V]{}
	rec.val = fn(key)
	rec.key = key
	return rec
}

func (rec *sonar[K, V]) Lookup(fn resume[K, V], key *K) Sonar[K, V] {
	rec.val = fn(key)
	rec.key = key
	return rec
}

func (rec *sonar[K, V]) eval() *V {
	if rec.val != nil {
		return rec.val
	}

	var backloop []writer[K, V]

	for _, p := range rec.chain {
		backloop = append(backloop, p.writer)
		if p.resume == nil {
			continue
		}
		if val := p.resume(rec.key); val != nil {
			rec.val = val
			break
		}
	}

	if rec.val != nil {
		for _, backing := range backloop {
			if backing == nil {
				continue
			}
			backing(rec.key, rec.val)
		}
	}

	return rec.val
}

func (rec *sonar[K, V]) Backing(rfn resume[K, V], wfn writer[K, V]) Sonar[K, V] {
	rec.chain = append(rec.chain, &pair[K, V]{rfn, wfn})
	return rec
}

func (rec *sonar[K, V]) MissResume(rfn resume[K, V]) Sonar[K, V] {
	return rec.Backing(rfn, nil)
}

func (rec *sonar[K, V]) WriteWith(wfn writer[K, V]) Sonar[K, V] {
	return rec.Backing(nil, wfn)
}

func (rec *sonar[K, V]) Get() *V {
	return rec.eval()
}

func (rec *sonar[K, V]) OrDft(dv *V) (*V, error) {
	eval := rec.eval()
	if eval != nil {
		return eval, nil
	}
	if dv != nil {
		return dv, nil
	}
	return nil, errors.New("default value is nil")
}
