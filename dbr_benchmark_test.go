package db_sql_benchmark

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	dbrDialect "github.com/gocraft/dbr/v2/dialect"
)

//
// Select benchmarks
//
func dbrToSQL(b dbr.Builder) (query string, args []interface{}) {
	// As ToSql method seems to be dropped, we use a trimmed version
	// of interpolator.encodePlaceholder method dbr calls under the hood.
	pbuf := dbr.NewBuffer()
	b.Build(dbrDialect.SQLite3, pbuf)
	return pbuf.String(), pbuf.Value()
}

func BenchmarkDbrSelectSimple(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dbrToSQL(dbr.Select("id").
			From("tickets").
			Where(
				dbr.And(
					dbr.Eq("subdomain_id", 1),
					dbr.Or(dbr.Eq("state", "open"), dbr.Eq("state", "spam")))))
	}
}

func BenchmarkDbrSelectSimpleRawWhere(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dbrToSQL(dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam"))
	}
}

func BenchmarkDbrSelectConditional(b *testing.B) {

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		qb := dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		if n%2 == 0 {
			qb.GroupBy("subdomain_id").
				Having("number = ?", 1).
				OrderBy("state").
				Limit(7).
				Offset(8)
		}

		dbrToSQL(qb)
	}
}
func BenchmarkDbrSelectComplex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dbrToSQL(dbr.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			// Where(dbr.Eq{"h": []int{1, 2, 3}}).
			GroupBy("i").
			GroupBy("ii").
			GroupBy("iii").
			Having("j = k").
			Having("jj = ?", 1).
			Having("jjj = ?", 2).
			OrderBy("l").
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8))
	}
}

func BenchmarkDbrSelectSubquery(b *testing.B) {

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		subQuery, _ := dbrToSQL(dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam"))

		dbrToSQL(dbr.Select("a", "b", fmt.Sprintf("(%s) AS subq", subQuery)).
			From("c").
			Distinct().
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8))
	}
}

//
// Insert benchmark
//
func BenchmarkDbrInsert(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dbr.InsertInto("mytable").
			Columns("id", "a", "b", "price", "created", "updated").
			Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05"))
	}
}

//
// Update benchmark
//
func BenchmarkDbrUpdateSetColumns(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dbr.Update("mytable").
			Set("foo", 1).
			Set("bar", dbr.Expr("COALESCE(bar, 0) + 1")).
			Set("c", 2).
			Where("id = ?", 9).
			Limit(10))
	}
}

func BenchmarkDbrUpdateSetMap(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dbr.Update("mytable").
			SetMap(map[string]interface{}{"b": 1, "c": 2, "bar": dbr.Expr("COALESCE(bar, 0) + 1")}).
			Where("id = ?", 9).
			Limit(10))
	}
}

//
// Delete benchmark
//
func BenchmarkDbrDelete(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dbr.DeleteFrom("test_table").
			Where("b = ?", 1).
			Limit(2))
	}
}
