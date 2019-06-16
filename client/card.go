package main

import (
	"fmt"
	"strconv"
)

// Card _
type Card struct {
	val int
	str string
}

// NewCard _
func NewCard(x int) *Card {
	if x < 11 {
		return &Card{x, strconv.Itoa(x)}
	}

	var str string

	switch x {
	case 11:
		str = "Jack"
	case 12:
		str = "Queen"
	case 13:
		str = "King"
	case 14:
		str = "Ace"
	default:
		panic("Invalid card value")
	}

	return &Card{x, str}
}

func (c *Card) isLower() bool {
	return c.val < 8
}

func (c *Card) isUpper() bool {
	return c.val > 7
}

// ToString _
func (c *Card) ToString() string {
	if c.isLower() {
		return fmt.Sprintf("%s (lower)", c.str)
	}

	return fmt.Sprintf("%s (upper)", c.str)
}
