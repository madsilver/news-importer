package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"silver.com/internal/infra/env"
)

type MysqlDB struct {
	Conn *sql.DB
}

func NewMysqlDB() *MysqlDB {
	conn, err := sql.Open("mysql", getDSN())
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("successfully connected to MySQL")

	return &MysqlDB{
		Conn: conn,
	}
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		env.GetString("MYSQL_USER", "silver"),
		env.GetString("MYSQL_PASSWORD", "silver"),
		env.GetString("MYSQL_HOST", "localhost"),
		env.GetString("MYSQL_PORT", "3306"),
		env.GetString("MYSQL_DATABASE", "db"))
}
