package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Errorf("Wrong cards number")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Suit: Spade, Rank: Ace}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades, got: ", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand should be
	// [40 35 ...]
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	shuffled := New(Shuffle)

	first := orig[40]
	second := orig[35]

	if shuffled[0] != first {
		t.Errorf("First card should be %v, but got %v", first, shuffled[0])
	}

	if shuffled[1] != second {
		t.Errorf("Second card should be %v, but got %v", second, shuffled[1])
	}

}

func TestJokers(t *testing.T) {
	numOfJokers := 3
	cards := New(Jokers(numOfJokers))
	count := 0
	for _, v := range cards {
		if v.Suit == Joker {
			count++
		}
	}

	if count != numOfJokers {
		t.Errorf("Expected %v Jokers, got %v", numOfJokers, count)
	}

}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, v := range cards {
		if v.Rank == Two || v.Rank == Three {
			t.Errorf("Got Two or Three in the Deck")
		}
	}

}

func TestDeck(t *testing.T) {
	numOfDecks := 3
	numOfCards := 13 * 4 * numOfDecks // 13 ranks, 4 suits
	cards := New(Deck(numOfDecks))

	if len(cards) != numOfCards {
		t.Errorf("Expected %v of cards, got %v", numOfCards, len(cards))
	}

}
