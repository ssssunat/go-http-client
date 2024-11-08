package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("happy path - can hit the api and return a pokemon", func(t *testing.T) {
		myClient := NewClient()
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})

	t.Run("sad path - return an error when the pokemon does not exist", func(t *testing.T) {
		myClient := NewClient()
		_, err := myClient.GetPokemonByName(context.Background(), "non-existant-pokemon")
		assert.Error(t, err)
	})
}
