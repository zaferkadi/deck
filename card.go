//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit type
type Suit uint8

// const sads
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // Special Card
)

var suites = [...]Suit{Spade, Diamond, Club, Heart}

// Rank type
type Rank uint8

const (
	_ Rank = iota
	Ace
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

// Card struct
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New is for initializing a deck with optional operation function
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	// iterate over the suits
	// iterate over the ranks
	// add suit, rank
	// return cards
	for _, suit := range suites {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// DefaultSort func
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort func
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less to sort
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Shuffle cards randomly
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))

	perm := shuffleRand.Perm(len(cards))

	for i, v := range perm {
		ret[i] = cards[v]
	}

	return ret
}

// Filter based on func filters
func Filter(f func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, v := range cards {
			if !f(v) {
				ret = append(ret, v)
			}
		}
		return ret
	}
}

// Jokers func with number of jokers
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}

// Deck func, create an N deck of cards
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card

		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
