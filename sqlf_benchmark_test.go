package db_sql_benchmark

import (
	"testing"

	"github.com/leporo/sqlf"
)

var s string

func sqlfSelectSimple(b *testing.B, dialect sqlf.Dialect) {
	for n := 0; n < b.N; n++ {
		q := dialect.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")
		s = q.String()
		q.Close()
	}
}

func sqlfSelectConditional(b *testing.B, dialect sqlf.Dialect) {
	for n := 0; n < b.N; n++ {
		q := dialect.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		if n%2 == 0 {
			q.GroupBy("subdomain_id").
				Having("number = ?", 1).
				OrderBy("state").
				Limit(7).
				Offset(8)
		}

		s = q.String()
		q.Close()
	}
}

func sqlfSelectComplex(b *testing.B, dialect sqlf.Dialect) {
	for n := 0; n < b.N; n++ {
		q := dialect.Select("DITINCT a, b, z, y, x").
			// Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where("g = ?", 3).
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
			Offset(8)
		s = q.String()
		q.Close()
	}
}

func sqlfSelectSubquery(b *testing.B, dialect sqlf.Dialect) {
	for n := 0; n < b.N; n++ {
		q := dialect.Select("DITINCT a, b").
			SubQuery("(", ") AS subq",
				sqlf.Select("id").
					From("tickets").
					Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")).
			From("c").
			// Distinct().
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where("g = ?", 3).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8)
		s = q.String()
		q.Close()
	}
}

func BenchmarkSqlfSelectSimple(b *testing.B) {
	sqlfSelectSimple(b, sqlf.NoDialect)
}

func BenchmarkSqlfSelectSimplePostgreSQL(b *testing.B) {
	sqlfSelectSimple(b, sqlf.PostgreSQL)
}

func BenchmarkSqlfSelectConditional(b *testing.B) {
	sqlfSelectConditional(b, sqlf.NoDialect)
}

func BenchmarkSqlfSelectConditionalPostgreSQL(b *testing.B) {
	sqlfSelectConditional(b, sqlf.PostgreSQL)
}

func BenchmarkSqlfSelectComplex(b *testing.B) {
	sqlfSelectComplex(b, sqlf.NoDialect)
}

func BenchmarkSqlfSelectComplexPostgreSQL(b *testing.B) {
	sqlfSelectComplex(b, sqlf.PostgreSQL)
}

func BenchmarkSqlfSelectSubquery(b *testing.B) {
	sqlfSelectSubquery(b, sqlf.NoDialect)
}

func BenchmarkSqlfSelectSubqueryPostgreSQL(b *testing.B) {
	sqlfSelectSubquery(b, sqlf.PostgreSQL)
}

//
// Insert benchmark
//
func sqlfInsert(b *testing.B, dialect sqlf.Dialect) {
	for n := 0; n < b.N; n++ {
		q := dialect.InsertInto("mytable").
			Set("id", 1).
			Set("a", "test_a").
			Set("b", "test_b").
			Set("price", 100.05).
			Set("created", "2014-01-05").
			Set("updated", "2015-01-05")
		s = q.String()
		q.Close()
	}
}

func BenchmarkSqlfInsert(b *testing.B) {
	sqlfInsert(b, sqlf.NoDialect)
}

func BenchmarkSqlfInsertPostgreSQL(b *testing.B) {
	sqlfInsert(b, sqlf.PostgreSQL)
}

//
// Update benchmark
//
func sqlfUpdateSetColumns(b *testing.B, dialect sqlf.Dialect) {
	for n := 0; n < b.N; n++ {
		q := dialect.Update("mytable").
			Set("foo", 1).
			SetExpr("bar", "COALESCE(bar, 0) + 1").
			Set("c", 2).
			Where("id = ?", 9).
			Limit(10)
		s = q.String()
		q.Close()
	}
}

func BenchmarkSqlfUpdateSetColumns(b *testing.B) {
	sqlfUpdateSetColumns(b, sqlf.NoDialect)
}

func BenchmarkSqlfUpdateSetColumnsPostgreSQL(b *testing.B) {
	sqlfUpdateSetColumns(b, sqlf.PostgreSQL)
}

//
// Delete benchmark
//
func sqlfDelete(b *testing.B, dialect sqlf.Dialect) {
	for n := 0; n < b.N; n++ {
		q := dialect.DeleteFrom("test_table").
			Where("b = ?", 1).
			Limit(2)
		s = q.String()
		q.Close()
	}
}

func BenchmarkSqlfDelete(b *testing.B) {
	sqlfDelete(b, sqlf.NoDialect)
}

func BenchmarkSqlfDeletePostgreSQL(b *testing.B) {
	sqlfDelete(b, sqlf.PostgreSQL)
}
