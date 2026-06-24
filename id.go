package nolfix

import "sync/atomic"

type IdGenerator interface {
	NextID() int
}

type IntGenerator struct {
	counter atomic.Int64
}

func (g *IntGenerator) NextID() int {
	return int(g.counter.Add(1))
}
