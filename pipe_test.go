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
			"pipe.ProcessAndConsume(p, c, %v) = %v, want %v",
			[]int{1, 2, 3}, actual, []int{0, 1, 2},
		)
	}
}

func TestProcess(t *testing.T) {
	p := pipe.New(pipe.StageFunc(func(x pipe.Item) pipe.Item {
		return x.(int) - 1
	}))
	actual := pipe.Process(p, 1, 2, 3)
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].(int) < actual[j].(int)
	})
	if actual[0] != 0 || actual[1] != 1 || actual[2] != 2 {
		t.Errorf(
			"pipe.Process(p, %v) = %v, want %v",
			[]int{1, 2, 3}, actual, []int{0, 1, 2},
		)
	}
}

func TestPipe(t *testing.T) {
	add := func(s string) func(pipe.Item) pipe.Item {
		return func(x pipe.Item) pipe.Item {
			return x.(string) + s
		}
	}
	p := pipe.New(
		pipe.StageFunc(add("a")),
		pipe.StageFunc(add("b")),
		pipe.StageFunc(add("c")),
	)
	p.Receive("")
	actual := p.Deliver()
	if actual != "abc" {
		t.Errorf("p.Deliver() = %s, want %s", actual, "abc")
	}
}

func TestDeliverThenReceive(t *testing.T) {
	p := pipe.New(pipe.StageFunc(func(x pipe.Item) pipe.Item {
		return x.(int) - 1
	}))
	p.Receive(1)
	actual := p.Deliver()
	if actual != 0 {
		t.Errorf("p.Deliver() = %d, want %d", actual, 0)
	}
	p.Receive(2)
	actual = p.Deliver()
	if actual != 1 {
		t.Errorf("p.Deliver() = %d, want %d", actual, 1)
	}
}

func TestNoStage(t *testing.T) {
	p := pipe.New()
	p.Receive(1)
	actual := p.Deliver()
	if actual != 1 {
		t.Errorf("p.Deliver() = %d, want %d", actual, 1)
	}
}
