package main

import (
	"os"
	"testing"
) // package for testing

func TestNewDeck(t *testing.T){ // initiate a test function 
	d:= newDeck() // create a new deck 

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d)) // test the len of the deck 

	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got  %v", d[0]) // to check is the first value is correct 
	}

	if d[len(d) -1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of clubs, but got  %v", d[len(d) -1]) // To check the last value
	}


}

// test save file
func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting") // remove file
	
	deck:=newDeck() // initiate a new deck 
	deck.saveToFile("_decktesting") // save deck in file

	loadedDeck := newDeckFromFile("_decktesting") // load file

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck)) // test len file
	}

	os.Remove("_decktesting") // remove file
}