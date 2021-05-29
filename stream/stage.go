package stream

import "stream_test/stream/types"

type Stage interface {
	Begin(int)
	Act(types.T)
	End()
	IsDone() bool
}

// impl

type stage struct {
	begin  func(int)
	act    func(types.T)
	end    func()
	isDone func() bool
}

func (s *stage) Begin(t int) {
	s.begin(t)
}

func (s *stage) Act(t types.T) {
	s.act(t)
}

func (s *stage) End() {
	s.end()
}

func (s *stage) IsDone() bool {
	return s.isDone()
}

type option func(*stage)

func WithBegin(begin func(int)) option {
	return func(s *stage) {
		s.begin = begin
	}
}

func WithAct(act func(types.T)) option {
	return func(s *stage) {
		s.act = act
	}
}

func WithEnd(end func()) option {
	return func(s *stage) {
		s.end = end
	}
}

func WithDoneCheck(check func() bool) option {
	return func(s *stage) {
		s.isDone = check
	}
}

func newStage(s Stage, opts ...option) *stage {
	t := &stage{
		begin:  s.Begin,
		act:    s.Act,
		end:    s.End,
		isDone: s.IsDone,
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func newTerminalStage(act func(types.T), opts ...option) *stage {
	t := &stage{
		begin:  func(int) {},
		act:    act,
		end:    func() {},
		isDone: func() bool { return false },
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}
