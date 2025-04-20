package db

import (
	"encoding/json"

	"TournamentProject/player"
)

type Db struct {
	players []*player.Player
}

func (db *Db) InMemPlayers() []*player.Player {
	return db.players
}

//

func (db *Db) StoreInMem(d *json.Decoder) error {
	var p []*player.Player
	if err := d.Decode(&p); err != nil {
		return err
	}
	db.players = p
	return nil
}
