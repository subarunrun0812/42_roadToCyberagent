package main


import (
	"encoding/json"
	"log"
	"net/http"
)

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

func main() {
	// ルーティング
	http.HandleFunc("/api/games", getGames)

	// サーバー起動
	port := "8000" // コンテナないのポート
	log.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
