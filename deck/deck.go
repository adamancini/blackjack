//go:generate stringer -type=Suit,Rank

// go install golang.org/x/tools/cmd/stringer@latest
// and the stringer package will generate String() code for the types
// specified

package deck

// create a Card type to be exported to represent a playing card

import (
	"fmt"
	"math/rand"
	"sort"
)

// `iota` starts at the first element and assigns a value of 0 and
// iterates until the last const
type Suit uint8

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

var suits = [...]Suit{Clubs, Diamonds, Hearts, Spades}

type Rank uint8

const (
	_   Rank = iota // would be nice for the string value to equate to a point value for some games
	Ace             // so skip the 0 index value with a _
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Rank
	Suit
}

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

func DefaultSort(c []Card) []Card {
	sort.Slice(c, Less(c))
	return c
}

func Sort(less func(c []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(c []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(c[i]) < absRank(c[j])
	}
}

func More(c []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(c[i]) > absRank(c[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Shuffle(r rand.Source) func([]Card) []Card {
	return func(c []Card) []Card {
		r := rand.New(r)
		ret := make([]Card, len(c))
		perm := r.Perm(len(c))
		for i, j := range r.Perm(len(c)) {
			ret[i] = c[j]
		}
		return ret
	}
	// ret := make([]Card, len(c))
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i, j := range r.Perm(len(c)) {
	// 	ret[i] = c[j]
	// }
	// return ret
}

func Filter(f func(c Card) bool) func([]Card) []Card {
	return func(c []Card) []Card {
		var ret []Card
		for _, card := range c {
			if !f(card) {
				ret = append(ret, card)
			}
		}
		return ret
	}
}

func Decks(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
