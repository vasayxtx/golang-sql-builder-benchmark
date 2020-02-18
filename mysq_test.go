package db_sql_benchmark

import (
	"database/sql"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

const mysqlDSN = "root:college@tcp(127.0.0.1:3306)/employees"

func makeMySQLDriver() *sql.DB {
	conn, _ := sql.Open("mysql", mysqlDSN)
	conn.SetMaxOpenConns(10)
	return conn
}

type Employee struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func TestDbrMySQL(t *testing.T) {
	sess := makeDbrSession()
	var emp Employee
	err := sess.Select("first_name", "last_name").
		From("employees").
		Where("emp_no = ?", 30000).
		LoadOne(&emp)
	require.NoError(t, err)
	require.Equal(t, "Matt", emp.FirstName)
	require.Equal(t, "Avouris", emp.LastName)
}

func TestSquirrelMySQL(t *testing.T) {
	driver := makeMySQLDriver()
	var emp Employee
	err := sq.Select("first_name", "last_name").
		From("employees").
		Where("emp_no = ?", 30000).
		RunWith(driver).QueryRow().Scan(&emp.FirstName, &emp.LastName)
	require.NoError(t, err)
	require.Equal(t, "Matt", emp.FirstName)
	require.Equal(t, "Avouris", emp.LastName)
}

func TestGoquMySQL(t *testing.T) {
	driver := makeMySQLDriver()
	db := goqu.New("mysql", driver)
	var emp Employee
	ok, err := db.Select("first_name", "last_name").
		From("employees").
		Where(goqu.L("emp_no = ?", 30000)).
		ScanStruct(&emp)
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, "Matt", emp.FirstName)
	require.Equal(t, "Avouris", emp.LastName)
}
