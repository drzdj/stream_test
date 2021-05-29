package stream

import "stream_test/stream/types"

func OfSlice(t interface{}) *stream {
	return &stream{
		it: newItSlice(t),
	}
}

func OfGenFunc(gen func() types.T) *stream {
	return &stream{
		it: newItGen(gen),
	}
}
