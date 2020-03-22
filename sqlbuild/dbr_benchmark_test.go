package sqlbuild

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	dbrDialect "github.com/gocraft/dbr/v2/dialect"
)

func dbrToSQL(b dbr.Builder) (query string, args []interface{}) {
	pbuf := dbr.NewBuffer()
	b.Build(dbrDialect.SQLite3, pbuf)
	return pbuf.String(), pbuf.Value()
}

func BenchmarkDbrSelectVerySimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qb := dbr.Select("id").
			From("tickets").
			Where(
				dbr.And(
					dbr.Eq("subdomain_id", 1),
					dbr.Or(
						dbr.Eq("state", "open"),
						dbr.Eq("state", "spam"),
					),
				),
			)
		dbrToSQL(qb)
	}
}

func BenchmarkDbrSelectVerySimpleRawExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qb := dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")
		dbrToSQL(qb)
	}
}

func BenchmarkDbrSelectSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qb := dbr.Select("id").
			From("tickets").
			Where(
				dbr.And(
					dbr.Eq("subdomain_id", 1),
					dbr.Or(
						dbr.Eq("state", "open"),
						dbr.Eq("state", "spam"),
					),
				),
			).
			GroupBy("subdomain_id").
			Having(dbr.Eq("number", 1)).
			OrderAsc("state").
			Limit(7).
			Offset(8)
		dbrToSQL(qb)
	}
}

func BenchmarkDbrSelectSimpleRawExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qb := dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			GroupBy("subdomain_id").
			Having("number = ?", 1).
			OrderAsc("state").
			Limit(7).
			Offset(8)
		dbrToSQL(qb)
	}
}

func BenchmarkDbrSelectComplex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qb := dbr.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where(
				dbr.And(
					dbr.Or(dbr.Eq("d", 1), dbr.Eq("e", "wat")),
					dbr.Eq("f", 2),
					dbr.Eq("x", "hi"),
					dbr.Eq("h", []int{1, 2, 3}),
				),
			).
			GroupBy("i", "ii", "iii").
			Having("j = k").
			Having(dbr.Eq("jj", 1)).
			Having(dbr.Eq("jjj", 2)).
			OrderAsc("l").
			OrderAsc("ll").
			OrderAsc("lll").
			Limit(7).
			Offset(8)
		dbrToSQL(qb)
	}
}
