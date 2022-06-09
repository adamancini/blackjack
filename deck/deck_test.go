package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCard(t *testing.T) {
	fmt.Println(Card{Rank: Ace, Suit: Hearts})

	// Output:
	// Ace of Hearts
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new standard deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	ex := Card{Rank: Ace, Suit: Clubs}
	if cards[0] != ex {
		t.Error("Did not get Ace of Clubs as first card. Got: ", cards[0])
	}
}

func TestCustomSort(t *testing.T) {
	cards := New(Sort(More))
	ex := Card{Rank: King, Suit: Spades}
	if cards[0] != ex {
		t.Error("Did not get King of Spades as first card. Got: ", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	cards := New()
	shuffled := Shuffle(cards)

	if reflect.DeepEqual(cards, shuffled) {
		t.Error("Shuffled deck is the same as reference deck")
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out, but got a two or a three")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Decks(3))

	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, but got %d cards", 13*4*3, len(cards))
	}
}
