//ff:func feature=repair type=command control=iteration dimension=1
//ff:what TestStartWorkers — every job gets an outcome and the channel closes when the jobs run out
package repair

import "testing"

func TestStartWorkers(t *testing.T) {
	db, c := loadFixture(t)
	jobs := make(chan Span, 3)
	for i, sp := range windowSpans(1, 5, 2) {
		sp.Seq = int64(i)
		jobs <- sp
	}
	close(jobs)
	seen := map[int64]bool{}
	for o := range startWorkers(db, c, jobs) {
		if o.Err != nil {
			t.Fatal(o.Err)
		}
		seen[o.Found.Seq] = true
	}
	if len(seen) != 3 {
		t.Fatalf("%v", seen)
	}
}
