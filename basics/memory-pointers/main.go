package main

import "fmt"

type BigData struct {
	//500 mb
}

// case 1	copies 500 mb data
func processBigData(bd BigData) {
	//bd -> copy of 500 mb in memory
}

// case 2 copies 5 bytes of memory addr
func processBigData2(bd *BigData) {
	//bd -> reference to original obj
}

type Player struct {
	health int
}

func (p *Player) takeDamagetFromExplosion() {
	p.health -= 10
}
func takeDamagetFromExplosion(player *Player) {
	fmt.Println("Player is taking damage from explosion")
	dmg := 10
	player.health -= dmg
}

func main() {
	// fmt.Println("Welcome to introduction to memory pointers")

	// var one int = 2
	// var ptr *int = &one

	// fmt.Println("value of pointer is (memory address)", ptr)

	// myNumber := 25

	// var addr = &myNumber
	// fmt.Println("value at pointer location is ", *addr)

	// *addr = *addr + 2
	// fmt.Println("value at pointer location after modification is ", *addr)

	player := Player{
		health: 100,
	}
	fmt.Println(player)
	//&player -> 8 byte long int
	takeDamagetFromExplosion(&player)
	fmt.Println(player) //health dosen't decrease
	player.takeDamagetFromExplosion()
	fmt.Println(player)

	takeDamagetFromExplosion(&player)
	fmt.Println(player)

}
