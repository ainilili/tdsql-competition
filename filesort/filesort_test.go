package filesort

import (
	"fmt"
	"github.com/ainilili/tdsql-competition/database"
	"github.com/ainilili/tdsql-competition/parser"
	"testing"
	"time"
)

func TestFileSorter_Sharding(t *testing.T) {
	db, _ := database.New("tdsqlshard-n756r9nq.sql.tencentcdb.com", 113, "nico", "Niconico2021@")
	tables, err := parser.ParseTables(db, "D:\\workspace-tencent\\data")
	if err != nil {
		t.Fatal(err)
	}
	fs, err := New(tables[0])
	if err != nil {
		t.Fatal(err)
	}
	s := time.Now().UnixNano()
	err = fs.Sharding()
	fmt.Println("sharding", (time.Now().UnixNano()-s)/1e6)
	if err != nil {
		t.Fatal(err)
	}

	s = time.Now().UnixNano()
	err = fs.Merging()
	fmt.Println("merging", (time.Now().UnixNano()-s)/1e6)
	if err != nil {
		t.Fatal(err)
	}
}
