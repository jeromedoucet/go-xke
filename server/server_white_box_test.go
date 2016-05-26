package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewServer(t *testing.T) {
	// given
	playerId := "player"
	bartenderUrl := "http://bartender.com"
	// when
	srv := NewServer(playerId, bartenderUrl)
	// then
	assert.Equal(t, playerId, srv.playerId)
	assert.Equal(t, bartenderUrl, srv.bartenderUrl)
}
