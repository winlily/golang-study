package demo

// 翻转带有中文数字字母的字符串
func ReverseString() string {
	str := "哈12测试XX下34Hh嘻"

	tmp := []rune(str)

	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	return string(tmp)
}
