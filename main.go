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

func (h Hand) MinScore() int {
	// assume aces are low
	score := 0

	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			// count only the first ace as an 11
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) Blackjack() bool {
	if len(h) == 2 {
		if h.Score() == 21 {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

gameloop:
	for input != "s" {
		if player.Blackjack() {
			fmt.Println("Blackjack. You Win!")
			break gameloop
		}

		if dealer.Blackjack() {
			fmt.Println("Dealer Blackjack.  You Lose!")
			break gameloop
		}
		fmt.Println("Player: ", player)
		fmt.Println("Dealer: ", dealer.DealerString())
		fmt.Println("What do you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			card, d = draw(d)
			player = append(player, card)
		case "s":
			break gameloop
		default:
			fmt.Println("That's not a valid option")
		}
	}

	pScore, dScore := player.Score(), dealer.Score()

	fmt.Println("** FINAL HANDS **")
	fmt.Println("Player: ", player, "\nScore: ", pScore)
	fmt.Println("Dealer: ", dealer, "\nScore: ", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted!")
	case dScore > 21:
		fmt.Println("The dealer busted!")
	case pScore > dScore:
		fmt.Println("You Win!")
	case dScore > pScore:
		fmt.Println("You Lose :(")
	case dScore == pScore:
		fmt.Println("Push")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
