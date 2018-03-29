package pipe_test

import (
	"fmt"

	"github.com/jwowillo/pipe"
)

func Example() {
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
	fmt.Println(pipe.Process(p, ""))
	// Output: [abc]
}
