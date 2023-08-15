package main

import "fmt"

type player struct {
	health int
	name string
}
// this returns an 8 byte pointer
// this pointer points to the spot in memory
// that holds a player struct
func createPlayer(name string) *player {
	const startHealth = 100
	p := player{name: name}
	p.health = startHealth
	return &p
}

func takeExplosionDmg(p *player) {
	p.health -= 10
}

func headshot(p *player) {
	p.health = 0
}

func main() {
	bradley := createPlayer("Bradley")
	fmt.Println(bradley)
	takeExplosionDmg(bradley)
	fmt.Println(bradley)
	takeExplosionDmg(bradley)
	fmt.Println(bradley)
	headshot(bradley)
	fmt.Println(bradley)
}

