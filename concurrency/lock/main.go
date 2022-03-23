package main

import (
	"fmt"
	"sync"
)

/*
使用一般的 waitGroup + Lock 效能會比較好。
不要因為使用 channel 比較潮，而強制在專案內使用。
在很多狀況底下一般的 Slice 或 callback 效能會比較好。
*/
func main() {
	foo := addByShareMemory(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}

func addByShareMemory(n int) []int {
	var ints []int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()
		}(i)
	}

	wg.Wait()
	return ints
}
