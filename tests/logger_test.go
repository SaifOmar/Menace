package tournament_test

import (
	"os"
	"testing"

	"TournamentProject/helpers"
)

func TestLogger(t *testing.T) {
	logger := helpers.NewTournamentLogger(true, "test.log")
	text := "test.test"
	logger.Info(text)
	f, err := os.ReadFile("../test.log")
	if err != nil {
		t.Error("can't create or write to disk files")
	}
	if string(f) != text {
		t.Error("write to disk files")
	}
}
