package pipe_test

import (
	"fmt"

	"github.com/jwowillo/pipe"
)

func Example_processAndConsume() {
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
	f := pipe.ConsumerFunc(func(x pipe.Item) {
		fmt.Println(x)
	})
	pipe.ProcessAndConsume(p, f, "")
	// Output:
	// abc
}
