package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"TournamentProject/db"
	"TournamentProject/tournament"
)

type server struct {
	*db.Db
}

func NewServer() *server {
	return &server{
		&db.Db{},
	}
}

func StartServer(t *tournament.Tournament) {
	server := NewServer()
	t.Db = server.Db
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(t.Matches)
	})
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		t.Run()
		json.NewEncoder(w).Encode(t.Matches)
	})

	http.HandleFunc("/newplayers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		err := server.StoreInMem(json.NewDecoder(r.Body))
		if err != nil {
			http.Error(w, "Failed to parse players", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Players stored successfully"})
		t.UseInMem()
		fmt.Println(server.InMemPlayers())
	})
	fmt.Println("Starting server on :8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// how will the data look like
// n players
// each player name
// each player attribute
// each player two abilities
// [[saif[attr[abilities]]]]
