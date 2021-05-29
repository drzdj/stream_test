package stream_test

import (
	"stream_test/stream"
	"stream_test/stream/types"
)

// - Q1: 输入一个整数 int，字符串string。将这个字符串重复n遍返回
func Question3Sub1(str string, n int) string {
	ret := make([]byte, 0, len(str)*n)
	stream.OfGenFunc(func() types.T {
		return str2bytes(str)
	}).Limit(n).Traverse(func(t types.T) {
		ret = append(ret, t.([]byte)...)
	})
	return bytes2str(ret)
}
