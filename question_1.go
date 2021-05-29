package stream_test

import (
	"stream_test/stream"
	"stream_test/stream/types"

	"github.com/Pallinder/go-randomdata"
)

// - Q1: 输入 employees，返回 年龄 >22岁 的所有员工，年龄总和
func Question1Sub1(employees []*Employee) int64 {
	ret := stream.OfSlice(employees).Peek(func(t types.T) {
	}).Filter(func(t types.T) bool {
		return *t.(*Employee).Age > 22
	}).ReduceWith(int64(0), func(r types.R, t types.T) types.R {
		return r.(int64) + int64(*t.(*Employee).Age)
	})
	return ret.(int64)
}

// - Q2: - 输入 employees，返回 id 最小的十个员工，按 id 升序排序
func Question1Sub2(employees []*Employee) []*Employee {
	n := 10
	ret := make([]*Employee, 0, n)
	stream.OfSlice(employees).Sort(func(l, r types.T) int {
		diff := l.(*Employee).Id - r.(*Employee).Id
		if diff == 0 {
			return 0
		} else if diff > 0 {
			return 1
		}
		return -1
	}).Limit(n).Traverse(func(t types.T) {
		ret = append(ret, t.(*Employee))
	})
	return ret
}

// - Q3: - 输入 employees，对于没有手机号为0的数据，随机填写一个
func Question1Sub3(employees []*Employee) []*Employee {
	return stream.OfSlice(employees).ReduceWith(make([]*Employee, 0, len(employees)), func(r types.R, t types.T) types.R {
		a := t.(*Employee)
		var ap string
		ap = *a.Phone
		if ap == "" {
			ap = randomdata.PhoneNumber()
		}
		tt := &Employee{
			Id:       a.Id,
			Name:     a.Name,
			Age:      a.Age,
			Position: a.Position,
			Phone:    &ap,
		}
		return append(r.([]*Employee), tt)
	}).([]*Employee)
}

// - Q4: - 输入 employees ，返回一个map[int][]int，其中 key 为 员工年龄 Age，value 为该年龄段员工ID
func Question1Sub4(employees []*Employee) map[int][]int64 {
	return stream.OfSlice(employees).ReduceWith(make(map[int][]int64, 100), func(r types.R, t types.T) types.R {
		rr := r.(map[int][]int64)
		tt := t.(*Employee)
		rr[*tt.Age] = append(rr[*tt.Age], tt.Id)
		return rr
	}).(map[int][]int64)
}
