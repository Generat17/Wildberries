package wq

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsync(t *testing.T) {
	wq := New(10)

	for i := 0; i < 20; i++ {
		i := i
		wq.AddTask(func(ctx context.Context) {
			log.Println("task #", i, "sleeping...")
			time.Sleep(time.Second)
			log.Println("task #", i, "done")
		})
	}

	start := time.Now()
	wq.RunAndWait(context.Background())
	delta := time.Since(start)

	log.Println("took", delta)
	assert.Less(t, delta, 2040*time.Millisecond)
}
