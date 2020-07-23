package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

type Crud interface {
	GetCreateQuery() string
	GetUpdateQuery() string
	GetDeleteQuery() string
}

func Create(data Crud) (uint64, error) {
	query, args, err := db.BindNamed(data.GetCreateQuery(), data)
	if err != nil {
		return 0, err
	}
	var id uint64
	err = db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Update(data Crud) error {
	return namedExec(data.GetUpdateQuery(), data)
}

func Delete(data Crud) error {
	return namedExec(data.GetDeleteQuery(), data)
}

func namedExec(query string, data interface{}) error {
	_, err := db.NamedExec(query, data)

	if err != nil {
		return err
	}

	return nil
}
