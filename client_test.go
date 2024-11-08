package client

import (
	"context"
	"net/http"
	"testing"
	"time"

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

	t.Run("happy path - testing the WithAPIURL option function", func(t *testing.T) {
		myClient := NewClient(
			WithAPIURL("my-test-url"),
		)
		assert.Equal(t, "my-test-url", myClient.apiURL)
	})

	t.Run("happy path - test with httpclient works", func(*testing.T) {
		myClient := NewClient(
			WithHTTPClient(&http.Client{
				Timeout: 1 * time.Second,
			}),
			WithAPIURL("my-test-url"),
		)

		assert.Equal(t, "my-test-url", myClient.apiURL)
		assert.Equal(t, 1 * time.Second, myClient.httpClient.Timeout)
	})
}
