package pipe_test

import (
	"fmt"

	"github.com/jwowillo/pipe"
)

func Example_produceProcessAndConsume() {
	p := pipe.New(
		pipe.StageFunc(func(x pipe.Item) pipe.Item {
			return x.(string) + "a"
		}),
		pipe.StageFunc(func(x pipe.Item) pipe.Item {
			return x.(string) + "b"
		}),
		pipe.StageFunc(func(x pipe.Item) pipe.Item {
			return x.(string) + "c"
		}),
	)
	var isDone bool
	pf := pipe.ProducerFunc(func() (pipe.Item, bool) {
		if isDone {
			return "", false
		}
		isDone = true
		return "", true
	})
	cf := pipe.ConsumerFunc(func(x pipe.Item) {
		fmt.Println(x)
	})
	pipe.ProduceProcessAndConsume(p, pf, cf)
	// Output:
	// abc
}
