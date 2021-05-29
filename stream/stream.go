package stream

import (
	"sort"
	"stream_test/stream/types"
)

type Stream interface {
	// stateless
	Filter(types.Predicate) Stream
	Map(types.Function) Stream
	Peek(types.Consumer) Stream

	// stateful
	Sort(types.Comparator) Stream
	Limit(int) Stream

	// terminative
	Traverse(types.Consumer)
	ReduceWith(types.R, types.Accumulator) types.R
	Count() int
}

// impl

type stream struct {
	it     Iterator
	wraps  []func(Stage) Stage
	fStage Stage
}

func (s *stream) calc() {
	st := s.fStage
	for i := len(s.wraps) - 1; i >= 0; i-- {
		st = s.wraps[i](st)
	}
	s.fStage = st
}

func (s *stream) terminate(fs Stage) {
	s.fStage = fs
	s.calc()
	stage, it := s.fStage, s.it
	stage.Begin(it.Size())
	for it.HasNext() && !stage.IsDone() {
		stage.Act(it.Next())
	}
	stage.End()
}

// stateless
func (s *stream) Filter(check types.Predicate) Stream {
	s.wraps = append(s.wraps, func(next Stage) Stage {
		return newStage(next, WithAct(func(t types.T) {
			if check(t) {
				next.Act(t)
			}
		}))
	})
	return s
}

func (s *stream) Map(mapf types.Function) Stream {
	s.wraps = append(s.wraps, func(next Stage) Stage {
		return newStage(next, WithAct(func(t types.T) {
			next.Act(mapf(t))
		}))
	})
	return s
}

func (s *stream) Peek(visit types.Consumer) Stream {
	s.wraps = append(s.wraps, func(next Stage) Stage {
		return newStage(next, WithAct(func(t types.T) {
			visit(t)
			next.Act(t)
		}))
	})
	return s
}

// stateful
func (s *stream) Sort(cmp types.Comparator) Stream {
	s.wraps = append(s.wraps, func(next Stage) Stage {
		var list []types.T
		return newStage(next,
			WithBegin(func(size int) {
				list = make([]types.T, 0, size)
				next.Begin(size)
			}), WithAct(func(t types.T) {
				list = append(list, t)
			}), WithEnd(func() {
				sort.Slice(list, func(i, j int) bool {
					return cmp(list[i], list[j]) < 0
				})
				next.Begin(len(list))
				it := newItSlice(list)
				for it.HasNext() {
					next.Act(it.Next())
				}
				next.End()
			}))
	})
	return s
}

func (s *stream) Limit(limit int) Stream {
	s.wraps = append(s.wraps, func(next Stage) Stage {
		cur := 0
		return newStage(next,
			WithAct(func(t types.T) {
				if cur >= limit {
					return
				}
				next.Act(t)
				cur++
			}),
			WithDoneCheck(func() bool {
				return cur >= limit
			}))
	})
	return s
}

// terminative
func (s *stream) Traverse(visit types.Consumer) {
	s.terminate(newTerminalStage(func(t types.T) {
		visit(t)
	}))

}

func (s *stream) ReduceWith(initVal types.R, acc types.Accumulator) types.R {
	ret := initVal
	s.terminate(newTerminalStage(func(t types.T) {
		ret = acc(ret, t)
	}))
	return ret
}

func (s *stream) Count() int {
	ret := 0
	s.terminate(newTerminalStage(func(t types.T) {
		ret++
	}))
	return ret
}
