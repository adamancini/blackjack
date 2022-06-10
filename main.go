package main

import (
	"fmt"

	"github.com/adamancini/blackjack/game"
)

func main() {
	var gs game.GameState
	gs = game.Shuffle(gs)

	for i := 0; i < 10; i++ {
		gs = game.Deal(gs)
		var input string
		for gs.State == game.StatePlayerTurn {
			fmt.Println("Player: ", gs.Player)
			fmt.Println("Dealer: ", gs.Dealer.DealerString())
			fmt.Println("What do you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = game.Hit(gs)
			case "s":
				gs = game.Stand(gs)
			default:
				fmt.Println("That's not a valid option")
			}
		}

		// dealer AI
		// if dealer score <= 16, hit
		// if dealer has a soft 17, hit
		for gs.State == game.StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = game.Hit(gs)
			} else {
				gs = game.Stand(gs)
			}
		}
		gs = game.EndHand(gs)
	}
}
