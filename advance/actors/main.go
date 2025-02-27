package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type Player struct {
	HP int
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
	case actor.Stopped:
		fmt.Println("Player actor is stopped (save state to db)")
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

	e.Send(pid, msg)

	//because e.Send is non blocking method
	time.Sleep(20 * time.Millisecond)

	fmt.Println("process id ", pid)

}
