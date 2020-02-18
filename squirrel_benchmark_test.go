package db_sql_benchmark

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
)

func BenchmarkSquirrelSelectVerySimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sq.Select("id").
			From("tickets").
			Where(
				sq.And{
					sq.Eq{"subdomain_id": 1},
					sq.Or{
						sq.Eq{"state": "open"},
						sq.Eq{"state": "spam"},
					},
				},
			).
			ToSql()
	}
}

func BenchmarkSquirrelSelectVerySimpleRawExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sq.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()
	}
}

func BenchmarkSquirrelSelectSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sq.Select("id").
			From("tickets").
			Where(
				sq.And{
					sq.Eq{"subdomain_id": 1},
					sq.Or{
						sq.Eq{"state": "open"},
						sq.Eq{"state": "spam"},
					},
				},
			).
			GroupBy("subdomain_id").
			Having(sq.Eq{"number": 1}).
			OrderBy("state").
			Limit(7).
			Offset(8).
			ToSql()
	}
}

func BenchmarkSquirrelSelectSimpleRawExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sq.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			GroupBy("subdomain_id").
			Having("number = ?", 1).
			OrderBy("state").
			Limit(7).
			Offset(8).
			ToSql()
	}
}

func BenchmarkSquirrelSelectComplex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sq.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where(
				sq.And{
					sq.Or{sq.Eq{"d": 1}, sq.Eq{"e": "wat"}},
					sq.Eq{"f": 2},
					sq.Eq{"x": "hi"},
					sq.Eq{"h": []int{1, 2, 3}},
				},
			).
			GroupBy("i", "ii", "iii").
			Having("j = k").
			Having(sq.Eq{"jj": 1}).
			Having(sq.Eq{"jjj": 2}).
			OrderBy("l").
			OrderBy("ll").
			OrderBy("ll").
			Limit(7).
			Offset(8).
			ToSql()
	}
}

func BenchmarkSquirrelRealMySQL(b *testing.B) {
	conn := makeMySQLDriver()
	var emp Employee
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sq.Select("first_name", "last_name").
			From("employees").
			Where("emp_no = ?", 30000).
			Limit(1).
			RunWith(conn).QueryRow().Scan(&emp.FirstName, &emp.LastName)
	}
}
