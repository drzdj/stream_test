package stream

import (
	"fmt"
	"reflect"
	"stream_test/stream/types"
)

type Iterator interface {
	HasNext() bool
	Next() types.T
	Size() int
}

type ItBase struct {
	cur int
	len int
}

func (it *ItBase) HasNext() bool {
	return it.cur < it.len
}

func (it *ItBase) Next() types.T {
	it.cur++
	return nil
}

func (it *ItBase) Size() int {
	return 0
}

type ItSlice struct {
	*ItBase
	data reflect.Value
}

func newItSlice(data interface{}) *ItSlice {
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		panic(fmt.Errorf("non slice input for data=%#v", data))
	}
	val := reflect.ValueOf(data)
	return &ItSlice{
		ItBase: &ItBase{
			cur: 0,
			len: val.Len(),
		},
		data: val,
	}
}

func (it *ItSlice) Next() types.T {
	ret := it.data.Index(it.cur).Interface()
	it.cur++
	return ret
}

func (it *ItSlice) Size() int {
	return it.len
}

type ItGen struct {
	*ItBase
	gen func() types.T
}

func newItGen(gen func() types.T) *ItGen {
	return &ItGen{
		ItBase: &ItBase{},
		gen:    gen,
	}
}

func (it *ItGen) Next() types.T {
	return it.gen()
}

func (it *ItGen) HasNext() bool {
	return true
}
