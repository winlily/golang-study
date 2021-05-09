package demo

// map 初始化及使用
func Map() {
	// 三步曲：申明+使用内置函数make进行初始化+赋值
	// 申明，此时m1是是未初始化的map，不可以进行赋值
	var m1 map[int]int
	// 使用make进行初始化
	m1 = make(map[int]int)
	// 赋值
	m1[1] = 1

	// 两步曲：Go短变量申明+赋值
	// 短变量申明
	m2 := make(map[int]int)
	// 赋值
	m2[1] = 11

	// 一步曲：初始化赋值一体
	m3 := map[int]int{
		1: 111,
	}
	m3[2] = 222
}
