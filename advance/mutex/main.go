package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Player struct {

	//adding mutex	// method 1
	// mu     sync.RWMutex
	// health int

	//method 2 - using atomic values
	health int32
}

func NewPlayer() *Player {
	return &Player{health: 100}
}

func (p *Player) getHealth() int {
	return int(atomic.LoadInt32(&p.health))
}
func (p *Player) setHealth(value int) {
	atomic.StoreInt32(&p.health, int32(value))
}

func startUILoop(p *Player) Player {
	ticker := time.NewTicker(time.Second)
	for {
		// p.mu.RLock()
		fmt.Printf("Player health %d\n", p.health)
		// p.mu.RUnlock()
		<-ticker.C

	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Microsecond * 300)
	for {
		// p.mu.Lock()
		// p.health -= rand.Int31()
		p.setHealth(rand.Intn(40))
		if p.health <= 0 {
			fmt.Println("Game over")
			break

		}
		// p.mu.Unlock()
		<-ticker.C

	}

}
func main() {

	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}
