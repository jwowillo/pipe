package pipe

import "sync"

// Consumer consumes Items received from Pipes.
type Consumer interface {
	Consume(Item)
}

// ConsumerFunc converts a function to a Consumer.
type ConsumerFunc func(Item)

// Consume applies the function to the Item.
func (f ConsumerFunc) Consume(x Item) {
	f(x)
}

// ProcessAndConsume processes all the Items with the Pipe and gives them to the
// Consumer as they are delivered.
func ProcessAndConsume(p *Pipe, c Consumer, xs ...Item) {
	for _, x := range xs {
		p.Receive(x)
	}
	var wg sync.WaitGroup
	wg.Add(len(xs))
	for i := 0; i < len(xs); i++ {
		go func() {
			c.Consume(p.Deliver())
			wg.Done()
		}()
	}
	wg.Wait()
}
