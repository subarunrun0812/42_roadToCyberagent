package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"path/to/your/project/dao"

	"github.com/google/uuid"
)

// DB接続用の変数
var db *sql.DB

// CreateUserRequestは/user/createで受け取るJSONの構造体
type CreateUserRequest struct {
	UsernName string `json:"username"`
	Coins     int    `json:"coins"`
	HighScore int    `json:"high_score"`
}

// Game 構造体
type Game struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ダミーデータ
var games = []Game{
	{ID: 1, Name: "Game 1"},
	{ID: 2, Name: "Game 2"},
}

// ゲーム一覧を返すハンドラー
func getGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(games)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// method check
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// ユーザーをDBに登録
	newUUID := uuid.New().String()
	user := dao.User{
		ID:        newUUID,
		UserName:  request.UsernName,
		Coins:     0, // coinsを0で初期化
		HighScore: 0, // highScoreを0で初期化
	}

	if err := dao.CreateUser(db, user); err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully", "id": newUUID})
}

func main() {
	// DSN(Data Source Name)をDocker Composeの設定に合わせる
	dsn := "user:pass@tcp(mysql:3306)/42tokyo-road-to-dojo-go_db-data?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	log.Println("Connected to database")

	// ルーティング
	http.HandleFunc("/api/games", getGames)
	http.HandleFunc("/user/create", createUser)

	// サーバー起動
	port := "8000"
	log.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
