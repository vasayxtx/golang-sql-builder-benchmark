package sqlbenchmark

import (
	"database/sql"
	"math/rand"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/stretchr/testify/require"
)

const mysqlDSN = "root:college@tcp(127.0.0.1:3306)/employees"

type Employee struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func BenchmarkMySQLRaw(b *testing.B) {
	db := openMySQL(false)
	defer db.Close()
	var emp Employee
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := loadEmployeeMySQLRaw(db, &emp, randomEmployeeID())
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMySQLRawWithoutPreparing(b *testing.B) {
	db := openMySQL(true)
	defer db.Close()
	var emp Employee
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := loadEmployeeMySQLRaw(db, &emp, randomEmployeeID())
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMySQLDbr(b *testing.B) {
	db := openMySQLDbr()
	defer db.Close()
	dbSess := db.NewSession(nil)
	var emp Employee
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := loadEmployeeDbr(dbSess, &emp, randomEmployeeID())
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMySQLGoquWithoutPreparing(b *testing.B) {
	db, sqlDB := openMySQLGoqu()
	defer sqlDB.Close()
	var emp Employee
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := loadEmployeeGoqu(db, &emp, randomEmployeeID(), false)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMySQLGoquWithPreparing(b *testing.B) {
	db, sqlDB := openMySQLGoqu()
	defer sqlDB.Close()
	var emp Employee
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := loadEmployeeGoqu(db, &emp, randomEmployeeID(), true)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMySQLDbr(t *testing.T) {
	db := openMySQLDbr()
	defer db.Close()
	dbSess := db.NewSession(nil)
	var emp Employee
	found, err := loadEmployeeDbr(dbSess, &emp, 30000)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, "Matt", emp.FirstName)
	require.Equal(t, "Avouris", emp.LastName)
}

func TestMySQLSquirrel(t *testing.T) {
	db := openMySQL(false)
	defer db.Close()

	var emp Employee
	err := sq.Select("first_name", "last_name").
		From("employees").
		Where("emp_no = ?", 30000).
		RunWith(db).QueryRow().Scan(&emp.FirstName, &emp.LastName)
	require.NoError(t, err)
	require.Equal(t, "Matt", emp.FirstName)
	require.Equal(t, "Avouris", emp.LastName)
}

func TestMySQLGoqu(t *testing.T) {
	db, sqlDB := openMySQLGoqu()
	defer sqlDB.Close()

	t.Run("prepare=true", func(t *testing.T) {
		var emp Employee
		found, err := loadEmployeeGoqu(db, &emp, 30000, true)
		require.NoError(t, err)
		require.True(t, found)
		require.Equal(t, "Matt", emp.FirstName)
		require.Equal(t, "Avouris", emp.LastName)
	})

	t.Run("prepare=true", func(t *testing.T) {
		var emp Employee
		found, err := loadEmployeeGoqu(db, &emp, 30000, false)
		require.NoError(t, err)
		require.True(t, found)
		require.Equal(t, "Matt", emp.FirstName)
		require.Equal(t, "Avouris", emp.LastName)
	})
}

func openMySQL(interpolateParams bool) *sql.DB {
	dsn := mysqlDSN
	if interpolateParams {
		dsn += "?interpolateParams=true"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(1)
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func openMySQLDbr() *dbr.Connection {
	db := openMySQL(false)
	return &dbr.Connection{DB: db, Dialect: dialect.MySQL, EventReceiver: &dbr.NullEventReceiver{}}
}

func openMySQLGoqu() (*goqu.Database, *sql.DB) {
	sqlDB := openMySQL(false)
	return goqu.New("mysql", sqlDB), sqlDB
}

func loadEmployeeMySQLRaw(db *sql.DB, emp *Employee, id int) (bool, error) {
	rows, err := db.Query("select first_name, last_name from employees where emp_no = ? limit 1", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if !rows.Next() {
		return false, nil
	}
	if err := rows.Scan(&emp.FirstName, &emp.LastName); err != nil {
		return false, err
	}
	return true, nil
}

func loadEmployeeDbr(dbSess dbr.SessionRunner, emp *Employee, id int) (bool, error) {
	err := dbSess.Select("first_name", "last_name").
		From("employees").
		Where("emp_no = ?", id).
		LoadOne(emp)
	if err != nil {
		if err == dbr.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func loadEmployeeGoqu(db *goqu.Database, emp *Employee, id int, prepare bool) (bool, error) {
	return db.Select("first_name", "last_name").
		From("employees").
		Where(goqu.Ex{"emp_no": id}).
		Prepared(prepare).
		ScanStruct(emp)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomEmployeeID() int {
	return random(10001, 499999)
}
