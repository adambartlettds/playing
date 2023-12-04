package main

import (
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	rand.Seed(now.UnixNano())
println("Numbers seeded using current date/time:", now.UnixNano())
	for i := 0; i < 5; i++ {
		println(rand.Intn(10))
	}
}
