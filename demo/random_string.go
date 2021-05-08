package demo

import (
	"math/rand"
	"time"
)

// 获取指定长度的随机字符串
// 场景：生成密码
func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)
	res := []byte{}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < n; i++ {
		res = append(res, bytes[r.Intn(len(bytes))])
	}

	return string(res)
}
