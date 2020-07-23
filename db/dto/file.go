package dto

import (
	"fmt"
	"github.com/schers/test-mafin/db"
)

const tableName string = "file"

type File struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
	Size int64  `db:"size"`
	Path string `db:"path"`
}

func (a *File) Save() error {

	if a.Id > 0 {
		err := db.Update(a)
		if err != nil {
			return err
		}
	} else {
		id, err := db.Create(a)
		if err != nil {
			return err
		}
		a.Id = id
	}

	return nil
}

func (a *File) Delete() error {
	err := db.Delete(a)
	if err != nil {
		return err
	}

	return nil
}

func (a *File) GetCreateQuery() string {
	return fmt.Sprintf(`INSERT INTO %s 
        (name, size, path) 
		VALUES 
		(:name, :size, :path)
		returning id`, tableName)
}

func (a *File) GetUpdateQuery() string {
	return fmt.Sprintf(`UPDATE %s SET 
     	name = :name, size = :size, path = :path WHERE id = :id`, tableName)
}

func (a *File) GetDeleteQuery() string {
	return fmt.Sprintf(`DELETE FROM %s WHERE id = :id`, tableName)
}
