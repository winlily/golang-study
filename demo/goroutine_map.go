package demo

import (
	"sync"
)

// 多协程并发数据竞争，以下一个协程在写入 m ，一个在读取 m，会产生数据竞争
// 可以使用 go run -race main.go 来检测数据竞争情况
func GoroutineMapErr() {
	m := make(map[int]int)

	go func() {
		m[1] = 1
	}()

	go func() {
		_ = m[1]
	}()

	select {}
}

// Go1.9以后版本，可以使用 sync.Map
func GoroutineMapSafety() {
	var m sync.Map

	go func() {
		m.Store(1, 1)
	}()

	go func() {
		m.Load(1)
	}()
}
