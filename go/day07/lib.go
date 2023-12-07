package main

import (
	"sort"
	"strconv"
	"strings"
)

var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type value string

func (v value) replace(o, n string) value {
	return value(strings.ReplaceAll(string(v), o, n))
}

func (v value) toString() string {
	return string(v)
}
func calculateHighestCombination(hand string, joker bool) (int, int) {
	hexa := value(hand).replace("T", "a").replace("J", "b").replace("Q", "c").replace("K", "d").replace("A", "e").toString()
	if joker {
		hexa = value(hexa).replace("b", "0").toString()
	}
	decimal64, _ := strconv.ParseInt(hexa, 16, 32)
	decimal := int(decimal64)
	cards := strings.Split(hand, "")
	sort.Slice(cards, func(i, j int) bool {
		return cardValues[cards[i]] > cardValues[cards[j]]
	})

	// Check for Five of a Kind.
	if cards[0] == cards[4] {
		// return "Five of a Kind"
		return 7, decimal
	}

	// Check for Four of a Kind.
	if cards[0] == cards[3] || cards[1] == cards[4] {
		// return "Four of a Kind"
		return 6, decimal
	}

	// Check for Full House.
	if (cards[0] == cards[2] && cards[3] == cards[4]) ||
		(cards[0] == cards[1] && cards[2] == cards[4]) {
		// return "Full House"
		return 5, decimal
	}

	// Check for Three of a Kind.
	if cards[0] == cards[2] || cards[1] == cards[3] || cards[2] == cards[4] {
		// return "Three of a Kind"
		return 4, decimal
	}

	// Check for Two Pair.
	if (cards[0] == cards[1] && cards[2] == cards[3]) ||
		(cards[0] == cards[1] && cards[3] == cards[4]) ||
		(cards[1] == cards[2] && cards[3] == cards[4]) {
		// return "Two Pair"
		return 3, decimal
	}

	// Check for One Pair.
	if cards[0] == cards[1] || cards[1] == cards[2] || cards[2] == cards[3] || cards[3] == cards[4] {
		// return "One Pair"
		return 2, decimal
	}

	// Default to High Card.
	// return "High Card"
	return 1, decimal
}
