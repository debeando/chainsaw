package main_test

import (
	"fmt"
	"testing"
	"os"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/swapbyt3s/chainsaw"
)

var db *sql.DB
var err error

func TestMain(t *testing.M) {
	db, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/")
	if err != nil {
			panic(err.Error())
	}
	defer db.Close()

	setupDatabase()

	code := t.Run()

	os.Exit(code)
}

func setupDatabase() {
	migrations := []string{
		"DROP DATABASE IF EXISTS demo_test;",
		"CREATE DATABASE IF NOT EXISTS demo_test;",
		"USE demo_test;",
		`CREATE TABLE IF NOT EXISTS demo_test.foo (
	id int NOT NULL auto_increment,
	value char(32) NOT NULL,
	token char(32),
	PRIMARY KEY (id)
);`,
	}

	for i := 1; i <= 100; i++ {
		migrations = append(migrations, fmt.Sprintf("INSERT INTO demo_test.foo (value) VALUES (MD5('%d'));", i))
	}

	for _, migration := range migrations {
		_, err := db.Query(migration)
		if err != nil {
			panic(err.Error())
		}
	}

}

//func TestSomeFeature(t *testing.T) {
//	rows, err := db.Query("SELECT id, value FROM demo_test.foo")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	for rows.Next() {
//		var id int
//		var val string
//
//		if err := rows.Scan(&id, &val); err != nil {
//			panic(err)
//		}
//
//		fmt.Printf("%d\t%s\n", id, val)
//	}
//}

func TestChunkCalculate(t *testing.T) {
	c := main.Chunk{Count: 1000, Delta: 100, Length: 100, Index: 550}
	c.Calculate()
	fmt.Printf("%d\n", c.Percentage)
	fmt.Printf("%d\n", c.Steps)
}
