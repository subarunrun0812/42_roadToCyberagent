package dao

import (
	"database/sql"
	"log"
)

type User struct {
	ID        string
	UserName  string
	Coins     int
	HighScore int
}

func CreateUser(db *sql.DB, user User) error {
	_, err := db.Exec("INSERT INTO users (id, username, coins, high_score) VALUES (?, ?, ?, ?)", user.ID, user.UserName, user.Coins, user.HighScore)
	if err != nil {
		log.Println("Error inserting user: ", err)
		return err
	}
	return nil
}
