//ff:func feature=repair type=command control=sequence
//ff:what TestWork — each batch's outcome carries its place, a found batch or the error of a failed one
package repair

import "testing"

func TestWork(t *testing.T) {
	db, c := loadFixture(t)
	jobs, out := make(chan Span, 2), make(chan outcome, 2)
	jobs <- Span{Seq: 0, From: 2, To: 3, Last: 3}
	close(jobs)
	work(db, c, jobs, out)
	if o := <-out; o.Err != nil || o.Found.Seq != 0 || o.Found.Last != 3 || o.Found.Entities != 2 {
		t.Fatalf("%+v", o)
	}
	jobs = make(chan Span, 1)
	jobs <- Span{Seq: 7, From: 2, To: 2, Last: 2}
	close(jobs)
	c.Schema += "_none"
	work(db, c, jobs, out)
	if o := <-out; o.Err == nil || o.Found.Seq != 7 {
		t.Fatalf("failed: %+v", o)
	}
}
