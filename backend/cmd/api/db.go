package main
// -代表就算沒用到以下的import你也要import來
import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

//假如未來某天想不只開一種db如+mongo, 可以讓connectToDB收到一個param並且在openDB裡面開另一個db
//這樣的話也要在repository/dbrepo底下平行加入一個mongo_dbrepo.go
func openDB(dsn string) (*sql.DB, error) {
	// driver name / connection string
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")
	return connection, nil
}