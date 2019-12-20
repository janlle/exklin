package exklin

import (
	"math/rand"
	"time"
)

func Random() (x int) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}
