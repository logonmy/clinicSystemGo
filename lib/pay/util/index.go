package pay

import (
	"fmt"
	"sort"
)

//SortArg 对象排序
func SortArg(para map[string]string) map[string]string {
	var keys []string
	var sortArg map[string]string
	for k := range para {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", para[k])
		sortArg[k] = para[k]
	}
	return sortArg
}

//CreateLinkstring 把所有元素，按照“参数=参数值”的模式用“&”字符拼接成字符串
func CreateLinkstring(para map[string]string) string {
	var ls string
	for k := range para {
		ls = ls + k + "=" + para[k] + "&"
	}
	ls = substr(ls, 0, 0)
	return ls
}

//substr 截取字符串 start 起点下标 length 需要截取的长度
func substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	if length == 0 {
		length = rl
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
