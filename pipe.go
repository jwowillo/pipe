// Package pipe allows Stages of functions to easily be assembled into concurent
// pipes.
package pipe

import (
	"sync"
)

// Item to be handled by a Stage.
type Item interface{}

// Stage handles an Item.
type Stage interface {
	Handle(Item) Item
}

// StageFunc converts a function to a Stage.
type StageFunc func(Item) Item

// Handle applies the function to the Item.
func (f StageFunc) Handle(x Item) Item {
	return f(x)
}

// Pipe connects Stages so many Items can be processed by the Stages in order
// concurrently.
type Pipe struct {
	m      sync.Mutex
	count  int
	stages []Stage
	links  []chan Item
}

// New Pipe with all the Stages connected in the order given.
func New(ss ...Stage) *Pipe {
	return &Pipe{
		m:      sync.Mutex{},
		stages: append([]Stage{}, ss...),
		links:  make([]chan Item, len(ss)+1),
	}
}

// Receive the Item into the beginning of the Pipe.
func (p *Pipe) Receive(x Item) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.isEmpty() {
		p.start()
	}
	p.count++
	go func() { p.links[0] <- x }()
}

// Deliver the item from the end of the Pipe once it's ready.
func (p *Pipe) Deliver() Item {
	p.m.Lock()
	defer p.m.Unlock()
	x := <-p.links[len(p.links)-1]
	p.count--
	if p.isEmpty() {
		p.stop()
	}
	return x
}

func (p *Pipe) isEmpty() bool {
	return p.count == 0
}

func (p *Pipe) start() {
	p.links[0] = make(chan Item)
	for i, stage := range p.stages {
		p.links[i+1] = make(chan Item)
		go func(receive <-chan Item, send chan<- Item, stage Stage) {
			for x := range receive {
				go func(x Item) {
					send <- stage.Handle(x)
				}(x)
			}
		}(p.links[i], p.links[i+1], stage)
	}
}

func (p *Pipe) stop() {
	for _, link := range p.links {
		close(link)
	}
}

// Process all the Items with the Pipe.
func Process(p *Pipe, xs ...Item) []Item {
	for _, x := range xs {
		p.Receive(x)
	}
	out := make([]Item, len(xs))
	for i := range out {
		out[i] = p.Deliver()
	}
	return out
}
