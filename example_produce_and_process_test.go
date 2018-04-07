package pipe_test

import (
	"fmt"

	"github.com/jwowillo/pipe"
)

func Example_produceAndProcess() {
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
	f := pipe.ProducerFunc(func() (pipe.Item, bool) {
		if isDone {
			return "", false
		}
		isDone = true
		return "", true
	})
	fmt.Println(pipe.ProduceAndProcess(p, f)[0])
	// Output:
	// abc
}
