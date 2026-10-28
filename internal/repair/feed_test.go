//ff:func feature=repair type=command control=sequence
//ff:what TestFeed — sends every span in order and closes; stops waiting for a slot or for a worker once stopped
package repair

import (
	"context"
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	spans := []Span{{Seq: 0}, {Seq: 1}}
	tokens, jobs := make(chan struct{}, 2), make(chan Span, 2)
	feed(context.Background(), spans, tokens, jobs)
	if a, b, open := <-jobs, <-jobs, false; a.Seq != 0 || b.Seq != 1 {
		t.Fatal("order")
	} else if _, open = <-jobs; open {
		t.Fatal("not closed")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	jobs = make(chan Span)
	feed(ctx, spans, make(chan struct{}), jobs)
	if _, open := <-jobs; open {
		t.Fatal("stopped waiting for a slot: not closed")
	}
	ctx, cancel = context.WithCancel(context.Background())
	time.AfterFunc(20*time.Millisecond, cancel)
	jobs = make(chan Span)
	feed(ctx, spans, make(chan struct{}, 1), jobs)
	if _, open := <-jobs; open {
		t.Fatal("stopped waiting for a worker: not closed")
	}
}
