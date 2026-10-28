//ff:func feature=repair type=writer control=sequence
//ff:what TestTakeLock — the lock is taken once; a second taker is refused until it is released; a closed handle fails
package repair

import (
	"context"
	"testing"
)

func TestTakeLock(t *testing.T) {
	db, schema := testSchema(t)
	c := Config{Schema: schema}
	ctx := context.Background()
	conn, err := takeLock(ctx, db, c)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := takeLock(ctx, db, c); err == nil {
		t.Fatal("second taker: want error")
	}
	if err := releaseLock(conn, c); err != nil {
		t.Fatal(err)
	}
	again, err := takeLock(ctx, db, c)
	if err != nil {
		t.Fatal("after release:", err)
	}
	releaseLock(again, c)
	if _, err := takeLock(ctx, closedDB(t), c); err == nil {
		t.Fatal("closed: want error")
	}
}
