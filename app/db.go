package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBService struct {
	writeDB *sqlx.DB
}

func NewDBService(c *Configuration) *DBService {
	create(c)
	writeDB, err := newDBFromConfig(c)
	if err != nil {
		panic(err)
	}
	return &DBService{
		writeDB: writeDB,
	}
}

func newDBFromConfig(c *Configuration) (*sqlx.DB, error) {
	con := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Address,
		c.Mysql.Port,
		c.Mysql.Database,
	)
	fmt.Println(con)
	db, err := sqlx.Connect(
		"mysql",
		con,
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func create(c *Configuration) {
	con := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Address,
		c.Mysql.Port,
	)
	fmt.Println(con)
	db, err := sql.Open("mysql", con)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS graphql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE graphql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user ( id integer, firstname varchar(255), lastname varchar(255) )")
	if err != nil {
		panic(err)
	}
	fmt.Println("")
}
