package pipe_test

import (
	"sort"
	"testing"

	"github.com/jwowillo/pipe"
)

func TestProduceAndProcess(t *testing.T) {
	p := pipe.New(pipe.StageFunc(func(x pipe.Item) pipe.Item {
		return x.(int) - 1
	}))
	list := []pipe.Item{1, 2, 3}
	i := 0
	f := pipe.ProducerFunc(func() (pipe.Item, bool) {
		if i >= len(list) {
			return nil, false
		}
		v := list[i]
		i++
		return v, true
	})
	actual := pipe.ProduceAndProcess(p, f)
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].(int) < actual[j].(int)
	})
	if actual[0] != 0 || actual[1] != 1 || actual[2] != 2 {
		t.Errorf(
			"pipe.ProduceAndProcess(p, f) = %v, want %v",
			actual, []int{0, 1, 2},
		)
	}
}

func TestProduceProcessAndConsume(t *testing.T) {
	p := pipe.New(pipe.StageFunc(func(x pipe.Item) pipe.Item {
		return x.(int) - 1
	}))
	list := []pipe.Item{1, 2, 3}
	i := 0
	pf := pipe.ProducerFunc(func() (pipe.Item, bool) {
		if i >= len(list) {
			return nil, false
		}
		v := list[i]
		i++
		return v, true
	})
	var actual []pipe.Item
	cf := pipe.ConsumerFunc(func(x pipe.Item) {
		actual = append(actual, x)
	})
	pipe.ProduceProcessAndConsume(p, pf, cf)
	sort.Slice(actual, func(i, j int) bool {
		return actual[i].(int) < actual[j].(int)
	})
	if actual[0] != 0 || actual[1] != 1 || actual[2] != 2 {
		t.Errorf(
			"pipe.ProduceProcessAndConsume(p, pf, cf) = %v, "+
				"want %v",
			actual, []int{0, 1, 2},
		)
	}
}
