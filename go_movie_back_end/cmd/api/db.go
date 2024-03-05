package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib" //go get github.com/jackc/pgx/v4/stdlib
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
//chatgpt instructions
type Authenticator struct {
    // Add fields related to authentication, such as secret keys, token expiration, etc.
    SecretKey string
    // Other fields...
}

