package main

import (
	"math/rand"
	"time"
)

func randomizer() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())
	return rand.New(seed)
}
