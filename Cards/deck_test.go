package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	var words []string
	for _, card := range cards {
		wordsInCard := strings.Split(card, " ")
		words = append(words, wordsInCard...)
	}
	suits := []string{"Hearts", "Spades", "Diamonds", "Cloves"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range suits {
		count := 0
		for _, word := range words {
			if word == suit {
				count++
			}
		}
		if count != 13 {
			t.Errorf("Expected 13 instances of "+suit+" got %v", count)
		}
	}
	for _, value := range values {
		count := 0
		for _, word := range words {
			if word == value {
				count++
			}
		}
		if count != 4 {
			t.Errorf("Expected 4 instances of "+value+" got %v", count)
		}
	}
	if len(cards) != 52 {
		t.Errorf("expected 52 cards but got %v", len(cards))
	}
}

func TestSaveToDeckAndReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	cards := newDeck()
	cards.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	if len(loadedDeck) != len(cards) {
		t.Errorf("expected %v cards but got %v.. might not be the correct deck", len(cards), len(loadedDeck))
	}
	os.Remove("_decktesting")
}
