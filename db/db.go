package db

import (
	"database/sql"
	"test/entity"
)

type DataBase struct {
	DB     *sql.DB
	DBName string
}

func NewDB(dbName string) (DataBase, error) {
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		return DataBase{}, err
	}
	return DataBase{
		DB:     db,
		DBName: dbName,
	}, nil
}
func (db *DataBase) CreateTable(ent entity.EntityInterface) error {
	db.DB.Query(ent.QueryCreateTable())
	return nil
}
