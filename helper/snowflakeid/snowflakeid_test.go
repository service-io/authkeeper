// Package snowflakeid
// @author tabuyos
// @since 2023/8/7
// @description snowflakeid
package snowflakeid

import (
	"fmt"
	"sync"
	"testing"
)

func TestGenerate(t *testing.T) {
	InitSnowflake()
	id := Generate()
	fmt.Printf("id: %d\n", id)
}

func TestLoopGenerate(t *testing.T) {
	InitSnowflake()
	c := 1000000
	wg := sync.WaitGroup{}
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			id := Generate()
			fmt.Printf("id: %d\n", id)
			wg.Done()
		}()
	}
	wg.Wait()
}
