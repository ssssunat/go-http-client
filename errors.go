package client

import (
	"errors"
	"fmt"
)

var (
	ErrFetchingPokemon = errors.New("failed to fetch pokemon")
)

type PokenonFetchErr struct {
	StatusCode int
	Message    string
}

func (e PokenonFetchErr) Error() string {
	return fmt.Sprintf("failed to fetch pokemon: %s with statuscode: %d", e.Message, e.StatusCode)
}
