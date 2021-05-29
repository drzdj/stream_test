package stream_test

import (
	"stream_test/stream"
	"stream_test/stream/types"
	"unsafe"
)

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// - Q1: 计算一个 string 中小写字母的个数
func Question2Sub1(str string) int64 {
	bytes := str2bytes(str)
	ret := stream.OfSlice(bytes).Filter(func(t types.T) bool {
		return t.(byte) >= 'a' && t.(byte) <= 'z'
	}).Count()
	return int64(ret)
}

// - Q2: 找出 []string 中，包含小写字母最多的字符串
func Question2Sub2(list []string) string {
	ret, cnt := "", int64(0)
	stream.OfSlice(list).Traverse(func(t types.T) {
		if tmp := Question2Sub1(t.(string)); cnt < tmp {
			ret, cnt = t.(string), tmp
		}
	})
	return ret
}
