package pipe_test

import (
	"sort"
	"testing"

	"github.com/jwowillo/pipe"
)

func TestProcessAndConsume(t *testing.T) {
	p := pipe.New(pipe.StageFunc(func(x pipe.Item) pipe.Item {
		return x.(int) - 1
	}))
	var actual []pipe.Item
	f := pipe.ConsumerFunc(func(x pipe.Item) {
		actual = append(actual, x)
	})
	pipe.ProcessAndConsume(p, f, 1, 2, 3)
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].(int) < actual[j].(int)
	})
	if actual[0] != 0 || actual[1] != 1 || actual[2] != 2 {
		t.Errorf(
			"pipe.ProcessAndConsume(p, f, %v) = %v, want %v",
			[]int{1, 2, 3}, actual, []int{0, 1, 2},
		)
	}
}
