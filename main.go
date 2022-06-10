package main

import (
	"fmt"

	"github.com/adamancini/blackjack/game"
)

func main() {
	var gs game.GameState
	gs = game.Shuffle(gs)

	for i := 0; i==; i < 10 {

	}
	gs = game.Deal(gs)

	// d := deck.New(deck.Decks(3), deck.Shuffle)

	// var card deck.Card
	// var player, dealer game.Hand

	// for i := 0; i < 2; i++ {
	// 	for _, hand := range []*game.Hand{&player, &dealer} {
	// 		card, d = deck.Draw(d)
	// 		*hand = append(*hand, card)
	// 	}
	// }

	var input string

gameloop:
	for gs.State == game.StatePlayerTurn {
		// if player.Blackjack() {
		// 	fmt.Println("Blackjack. You Win!")
		// 	break gameloop
		// }

		// if dealer.Blackjack() {
		// 	fmt.Println("Dealer Blackjack.  You Lose!")
		// 	break gameloop
		// }
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

	// change these to store in variables so we don't have to recalculate  every single time

	// for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
	// 	card, d = deck.Draw(d)
	// 	dealer = append(dealer, card)
	// }

	gs = game.EndHand(gs)
}
