//ff:func feature=repair type=writer control=sequence
//ff:what TestReleaseLock — releasing frees the lock for another taker; releasing a closed connection fails
package repair

import (
	"context"
	"testing"
)

func TestReleaseLock(t *testing.T) {
	db, schema := testSchema(t)
	c := Config{Schema: schema}
	conn, err := takeLock(context.Background(), db, c)
	if err != nil {
		t.Fatal(err)
	}
	if err := releaseLock(conn, c); err != nil {
		t.Fatal(err)
	}
	if err := releaseLock(conn, c); err == nil {
		t.Fatal("closed connection: want error")
	}
	next, err := takeLock(context.Background(), db, c)
	if err != nil {
		t.Fatal("free again:", err)
	}
	releaseLock(next, c)
}
