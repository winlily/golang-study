package demo

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
