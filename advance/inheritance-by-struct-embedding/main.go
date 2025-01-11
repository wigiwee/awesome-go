package main

import "fmt"

type SpecialPosition struct {
	Position
}

func (s *SpecialPosition) MoveSpecial(x, y float64) {
	s.x += x * x
	s.y += y * y
}

type Position struct {
	x float64
	y float64
}

// normal way to do
type Player struct {
	// posX float64
	// posY float64
	*Position
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

// func (p *Player) Move(x, y float64) {
// 	p.posX += x
// 	p.posY += y
// }

// func (p *Player) Teleport(x, y float64) {
// 	p.posX = x
// 	p.posY = y
// }

// type Enemy struct {
// 	posX float64
// 	posY float64
// }

// func (e *Enemy) Move(x, y float64) {
// 	e.posX += x
// 	e.posY += y
// }

// func (e *Enemy) Teleport(x, y float64) {
// 	e.posX = x
// 	e.posY = y
// }

// type Enemy struct {
// 	*Position
// }

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{
			// Position: &Position{},
		},
	}
}

func main() {
	Player := NewPlayer()
	fmt.Println(Player.Position)

	Player.Move(2.5, 3.5)

	fmt.Println(Player.Position)
	Player.Teleport(65, 22)
	fmt.Println(Player.Position)

	raidBoss := NewEnemy()
	fmt.Println("RaidBoss pos", raidBoss.Position)
	raidBoss.Move(3, 6)
	fmt.Println("RaidBoss pos", raidBoss.Position)
	raidBoss.MoveSpecial(3, 5)
	fmt.Println("RaidBoss pos", raidBoss.Position)

}
