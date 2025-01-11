package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	kickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}
type Messi struct {
	stamina  int
	power    int
	charisma int
}

func (m Messi) kickBall() {
	shot := m.stamina + m.power*m.charisma
	fmt.Println("Messi kicked the ball,", shot)
}

func (f FootballPlayer) kickBall() {
	shot := f.stamina + f.power
	fmt.Println("I am kickint the ball,", shot)
}

// func main() {
// 	team := make([]FootballPlayer, 11)
// 	for i := 0; i < len(team); i++ {
// 		team[i] = FootballPlayer{
// 			stamina: rand.Intn(10),
// 			power:   rand.Intn(10),
// 		}
// 	}
// 	for i := 0; i < len(team); i++ {
// 		team[i].kickBall()
// 	}
// }

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-1; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[10] = Messi{
		stamina:  10,
		power:    10,
		charisma: 8,
	}
	for i := 0; i < len(team); i++ {
		team[i].kickBall()
	}
}
