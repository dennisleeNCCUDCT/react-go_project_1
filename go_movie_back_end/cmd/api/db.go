package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
)

func openDB(dsn string) (*sql.DB, error){
	db, err :=sql.Open("pgx", dsn)
	if err != nil{
		return nil, err
	}
	err=db.Ping()
	if err != nil{
		return nil, err
	}

	return db ,nil
}

func (app *application) connectDB() (*sql.DB,error){
connection, err :=openDB(app.DSN)
if err!=nil{
	return nil, err
}

log.Println("Connect to Postgres")
return connection, nil

}
