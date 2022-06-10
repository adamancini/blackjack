package game

import (
	"fmt"
	"strings"

	"github.com/adamancini/blackjack/deck"
)

// Deck: [A, 10, J, ...]
// Turn: NewGame
// Player Hand: []
// Dealer Hand: []ret.Player

// Deck: [3, 6, Q, ...]
// Turn: HandsDealt
// Player Hand: [A, J]
// Dealer Hand: [10, 8]

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

type GameState struct {
	Deck   []deck.Card
	Turn   int
	State  State
	Player Hand // remember that a Hand is a slice in the implementation
	Dealer Hand
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

// func Deal(gs GameState) GameState {
// 	ret := clone(gs)
// 	... ret
// 	return ret
// }

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		Turn:   gs.Turn,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Decks(3), deck.Shuffle)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 6)
	ret.Dealer = make(Hand, 0, 6)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = deck.Draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = deck.Draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}

}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(gs)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()

	fmt.Println("** FINAL HANDS **")
	fmt.Println("Player: ", ret.Player, "\nScore: ", pScore)
	fmt.Println("Dealer: ", ret.Dealer, "\nScore: ", dScore)
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
	fmt.Println()

	ret.Player = nil
	ret.Dealer = nil
	return ret
}
