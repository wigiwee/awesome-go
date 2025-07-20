package main

import (
	"fmt"
	"log"

	"github.com/anthdm/hollywood/actor"
)

type Player struct {
	HP           int
	inventoryPid *actor.PID
}

type Inventory struct {
	Bottles int
}

type DrinkBottle struct {
	amount int
}

func NewInventory(bottles int) actor.Producer {
	return func() actor.Receiver {
		return &Inventory{
			Bottles: bottles,
		}
	}
}

func (i *Inventory) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("Inventory actor is started (read state from DB) ")
		_ = msg
	case *actor.Stopped:
		fmt.Println("Inventory actor stopped(save state to DB)")

	}
}

func NewPlayer(hp int) actor.Producer {
	// return &Player{}
	return func() actor.Receiver {
		return &Player{
			HP: hp,
		}
	}
}

type takeDamage struct {
	amount int
}

func (p *Player) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case takeDamage:
		fmt.Println("Player is taking damage ", msg.amount)
	case actor.Started:
		fmt.Println("player actor is started (read state from db)")
		p.inventoryPid = c.SpawnChild(NewInventory(1), "inventory")
	case actor.Stopped:
		fmt.Println("Player actor is stopped (save state to db)")
	case DrinkBottle:

	}
}

func main() {
	fmt.Println("getting started with actors")
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	pid := e.Spawn(NewPlayer(100), "player", actor.WithID("myUsedId69"))

	msg := takeDamage{amount: 333}

	_ = msg
	//e.Send(pid, msg)

	//because e.Send is non blocking method
	//time.Sleep(20 * time.Millisecond)

	fmt.Println("process id ", pid)

}
