package hand

import (
	//Imports card.go as c to prevent redundant card.Card or card.NewCard(int, string) every time a card is created
	c "example/web-service-gin/card"
	"math/rand"
	"time"
)

// Defining hand as an array of cards separate from the deck
type Hand struct {
	ActualHand []c.Card //the 5 cards in the hand
	HandType   string   //i.e. straight, 4 of a kind, royal flush...
}

type Getter interface {
	GetAll() []Card
}

type Adder interface {
	Add(card Card)
}

// array of cards inputted by user
type UserHand struct {
	Cards []Card
}

func New() *UserHand {
	return &UserHand{
		Cards: []Card{},
	}
}

func (r *UserHand) Add(card Card) {
	r.Cards = append(r.Cards, card)
}

func (r *UserHand) GetAll() []Card {
	return r.Cards
}

// Creating a new hand
func NewHand(handType string) *Hand {
	return &Hand{[]c.Card{}, handType}
}

// Checking the card at an index in the hand
func CheckCardIndex(hand *Hand, index int) c.Card {
	return hand.ActualHand[index]
}

// Add a specific card to the hand if hand has less than 5 cards ,for testing only
func AddCardHandSpecific(hand *Hand, val int, suit string) string {
	if val < 0 || val > 12 {
		return "adding specific card: value is invalid"
	} else if suit != "Heart" && suit != "Diamond" && suit != "Club" && suit != "Spade" {
		return "adding specific card: suit is invalid"
	} else {
		if len(hand.ActualHand) < 5 {
			hand.ActualHand = append(hand.ActualHand, c.NewCard(val, suit))
			return "adding specific card: successful"
		} else {
			return "adding specific card: length of hand is already 5"
		}
	}
}

// Add a random card to the hand if hand has less than 5 cards
func AddCardHandRandom(hand *Hand) string {
	if len(hand.ActualHand) < 5 {
		//get "random" value from time
		rand.Seed(time.Now().UnixNano())
		Number := rand.Intn(13)
		SuitValue := rand.Intn(4)
		Suit := ""
		switch SuitValue {
		case 0:
			Suit = "Heart"
		case 1:
			Suit = "Diamond"
		case 2:
			Suit = "Club"
		case 3:
			Suit = "Spade"
		default:
			Suit = "Error when adding card to hand" //this should not ever be the value so if it is then an error has occured
		}
		card := c.NewCard(Number, Suit)
		hand.ActualHand = append(hand.ActualHand, card)
		return "adding random card: successful"
	} else {
		return "adding random card: length of hand is already 5"
	}
}
