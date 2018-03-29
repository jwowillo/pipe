package pipe_test

import (
	"testing"
	"time"

	"github.com/jwowillo/pipe"
)

// Manual test to verify concurrency. Should finish in around 1.5 seconds since
// the longest word has 5 characters and each character causes a 0.1 second
// pause and there are 3 stages.
func BenchmarkPipe(b *testing.B) {
	f := pipe.StageFunc(func(x pipe.Item) pipe.Item {
		for _ = range x.(string) {
			time.Sleep(time.Second / 10)
		}
		return x
	})
	p := pipe.New(f, f, f)
	p.Receive("cat")
	p.Receive("wolf")
	p.Receive("mouse")
	p.Deliver()
	p.Deliver()
	p.Deliver()
}
