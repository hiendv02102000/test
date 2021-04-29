package main

import (
	"fmt"
	"math/rand"
	"test/demo_goroutine_channel"
)

func Battle(c, quit chan bool, b1, b2 *demo_goroutine_channel.Bot) {
	for {
		select {
		case <-quit:
			{
				switch {
				case b1.Health == b2.Health:
					{
						fmt.Println("Tie")
					}
				case b1.Health > b2.Health:
					{
						fmt.Println("Bot1 Win")
					}
				case b1.Health < b2.Health:
					{
						fmt.Println("Bot2 Win")
					}
				}

				return
			}
		case turn_bot1 := <-c:
			{
				if turn_bot1 {
					b1.Attack(b2)
				} else {
					b2.Attack(b1)
				}
			}

		}
		fmt.Printf("Bot1:%d >>><<< Bot2:%d\n", b1.Health, b2.Health)
	}
}
func main() {
	bot1 := demo_goroutine_channel.Bot{
		Health: 100,
		Damage: 5,
	}
	bot2 := demo_goroutine_channel.Bot{
		Health: 100,
		Damage: 4,
	}
	turn := make(chan bool)
	quit := make(chan bool)
	go func() {

		for bot1.Health > 0 && bot2.Health > 0 {

			turn_bot1 := (rand.Int()%2 == 0)
			turn <- turn_bot1

		}
		quit <- true

	}()
	Battle(turn, quit, &bot1, &bot2)

}
