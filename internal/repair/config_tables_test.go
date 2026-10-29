//ff:func feature=repair type=helper control=sequence
//ff:what TestConfigTables — an order build uses reference_snak_order.sql and its checkpoint table, a repair the repair work tables
package repair

import (
	"strings"
	"testing"
)

func TestConfigTables(t *testing.T) {
	ddl, cp := Config{Order: true}.tables()
	if cp != "reference_snak_order_checkpoint" || !strings.Contains(ddl, "CREATE TABLE IF NOT EXISTS reference_snak_order (") || strings.Contains(ddl, "repair_reference_snak_done") {
		t.Fatalf("order: %s %s", cp, ddl)
	}
	ddl, cp = Config{}.tables()
	if cp != "repair_reference_snak_checkpoint" || !strings.Contains(ddl, "repair_reference_snak_done") {
		t.Fatalf("repair: %s %s", cp, ddl)
	}
}
