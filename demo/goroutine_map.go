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

// Go1.9以后版本，可以使用 sync.Map，只对读多写少的场景有效率提升
func GoroutineMapSafety1() {
	var m sync.Map

	for i := 0; i < 10000; i++ {
		go func() {
			m.Store(1, 1)
		}()

		go func() {
			m.Load(1)
		}()
	}
}

var lock sync.RWMutex

// 使用读写锁
func GoroutineMapSafety2() {
	m := make(map[int]int)

	for i := 0; i < 10000; i++ {
		go func() {
			lock.Lock()
			m[1] = 1
			lock.Unlock()
		}()

		go func() {
			lock.RLock()
			_ = m[1]
			lock.RUnlock()
		}()
	}
}
