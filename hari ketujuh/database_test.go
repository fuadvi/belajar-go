package hari_ketujuh

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_database_golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestExecSql(t *testing.T) {
	db := GetConnectionDb()
	defer db.Close()

	ctx := context.Background()
	command := "INSERT INTO coustumer (id, name) VALUES (2,'maulana')"

	_, err := db.ExecContext(ctx, command)
	if err != nil {
		panic(err)
	}

	fmt.Println("berhasil create database")
}

func TestQuerySql(t *testing.T) {
	db := GetConnectionDb()
	defer db.Close()

	ctx := context.Background()
	command := "SELECT * FROM coustumer"

	rows, err := db.QueryContext(ctx, command)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)
		fmt.Println("name :", name)
	}
	defer rows.Close()
}

func TestQueryComplex(t *testing.T) {
	db := GetConnectionDb()
	defer db.Close()

	ctx := context.Background()

	command := "SELECT id, name, email, belance, rating, bird_date, married, create_at FROM coustumer"

	rows, err := db.QueryContext(ctx, command)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, balance int
		var name string
		var email sql.NullString
		var rating float32
		var brithDate, createAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &brithDate, &married, &createAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)
		fmt.Println("name :", name)
		fmt.Println("email :", email.String)
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		fmt.Println("brith_date :", brithDate)
		fmt.Println("married :", married)
		fmt.Println("create_at :", createAt)
	}

	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnectionDb()
	defer db.Close()

	ctx := context.Background()

	var username = "admin"
	var password = "admin"
	QuerySql := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	row, err := db.QueryContext(ctx, QuerySql, username, password)
	if err != nil {
		panic(err)
	}

	if row.Next() {
		var username string
		err := row.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("berhasil login")
	} else {
		fmt.Println("gagal login")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnectionDb()
	defer db.Close()

	ctx := context.Background()

	command := "INSERT INTO comments (email, commentar) values (?, ?)"
	result, err := db.ExecContext(ctx, command, "joe@gmail.com", "ini komentar joe")
	if err != nil {
		panic(err)
	}

	LastIndex, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println(LastIndex)
}
