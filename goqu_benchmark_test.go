package db_sql_benchmark

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
)

var driver *sql.DB

func init() {
	db, _, _ := sqlmock.New()
	driver = db
}

func BenchmarkGoquSelectVerySimple(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		db.From("tickets").
			Where(
				goqu.And(
					goqu.I("subdomain_id").Eq(1),
					goqu.Or(
						goqu.I("state").Eq("open"),
						goqu.I("state").Eq("spam"),
					),
				),
			).
			ToSQL()
	}
}

func BenchmarkGoquSelectVerySimpleRawExp(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		db.From("tickets").
			Where(goqu.L("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")).
			ToSQL()
	}
}

func BenchmarkGoquSelectSimple(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		db.Select("id").
			From("tickets").
			Where(
				goqu.And(
					goqu.I("subdomain_id").Eq(1),
					goqu.Or(
						goqu.I("state").Eq("open"),
						goqu.I("state").Eq("spam"),
					),
				),
			).
			GroupBy("subdomain_id").
			Having(goqu.I("number").Eq(1)).
			Order(goqu.I("state").Asc()).
			Limit(7).
			Offset(8).
			ToSQL()
	}
}

func BenchmarkGoquSelectSimpleRawExp(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		db.Select("id").
			From("tickets").
			Where(goqu.L("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")).
			GroupBy("subdomain_id").
			Having(goqu.L("number = ?", 1)).
			Order(goqu.I("state").Asc()).
			Limit(7).
			Offset(8).
			ToSQL()
	}
}

func BenchmarkGoquSelectComplex(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		db.Select(goqu.DISTINCT("a"), "b", "z", "y", "x").
			From("c").
			Where(
				goqu.Or(goqu.I("d").Eq(1), goqu.I("e").Eq("wat")),
				goqu.I("f").Eq(2),
				goqu.I("x").Eq("hi"),
				goqu.I("h").Eq([]int{1, 2, 3}),
			).
			GroupBy("i").
			GroupBy("ii").
			GroupBy("iii").
			Having(goqu.L("j = k")).
			Having(goqu.I("jj").Eq(1)).
			Having(goqu.I("jjj").Eq(2)).
			Order(goqu.I("l").Asc()).
			Order(goqu.I("ll").Asc()).
			Order(goqu.I("lll").Asc()).
			Limit(7).
			Offset(8).
			ToSQL()
	}
}
