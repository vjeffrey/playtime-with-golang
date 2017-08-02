package main

import (
	"database/sql"
	"fmt"
	"github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "vjeffrey"
	password = ""
	dbname   = "adventure-game"
)

const (
	MySQL = iota
	Postgres
	SQLite
)

var CREATE = map[int]string{
	Postgres: "CREATE TABLE IF NOT EXISTS playerdata(uid INTEGER PRIMARY KEY, username text DEFAULT 'user', score INTEGER DEFAULT 0, health INTEGER DEFAULT 0, created date DEFAULT '2012-08-04')",
}

func setUpDB(player Character) {
	logrus.Debugf("starting the db")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec("DROP TABLE IF EXISTS playerdata"); err != nil {
		panic(err)
	}

	if _, err = db.Exec(CREATE[Postgres]); err != nil {
		panic(err)
	}

	fmt.Println("# Inserting values")
	var lastInsertId int

	sqlStatement := `INSERT INTO playerdata VALUES ($1, $2, $3, $4)`

	_, err = db.Exec(sqlStatement, 7, player.Name, 100, player.Health)
	// err = db.QueryRow("INSERT INTO playerdata VALUES(1, 'name', 20, 80, '1971-07-13')").Scan(&lastInsertId)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
		} else {
			panic(err)
		}
	}
	fmt.Println("last inserted id =", lastInsertId)

	// r := rand.New(rand.NewSource(99))
	// sqlStatement := `INSERT INTO playerdata VALUES ($1, $2, $3, $4)`
	// _, err = db.Exec(sqlStatement, r, player.Name, 0, player.Health)

	rows, err := db.Query("SELECT * FROM playerdata")

	var uid int
	var username string
	var score int
	var health int
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &score, &health, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(score)
		fmt.Println(health)
		fmt.Println(created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
