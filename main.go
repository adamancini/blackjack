package main

import (
	"fmt"
	"strings"

	"github.com/adamancini/blackjack/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", HIDDEN"
}

func main() {
	d := deck.New(deck.Decks(3), deck.Shuffle)

	var card deck.Card
	var player, dealer Hand

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, d = draw(d)
			*hand = append(*hand, card)
		}
	}

	var input string
	for input != "s" {
		fmt.Println("Player: ", player)
		fmt.Println("Dealer: ", dealer.DealerString())
		fmt.Println("What do you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			card, d = draw(d)
			player = append(player, card)
		default:
			fmt.Println("That's not a valid option")
		}
	}

	fmt.Println("** FINAL HANDS **")
	fmt.Println("Player: ", player)
	fmt.Println("Dealer: ", dealer)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
