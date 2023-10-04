package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string // we initiate the type of 'deck' for every file with the variable deck 

func newDeck() deck{
	cards:= deck{}
    // one list
	cardSuits:=[]string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}
	// one list
	cardValues := []string{"Ace", "Two","three", "Four"}

	for _, suit := range cardSuits { //one loop +"_" to say to go we do not want use them
		for _, value := range cardValues{ // one loop inside the first loop + same 
			cards = append (cards, value+" of "+suit)
		}
	}
	return cards 
}
// to print deck 
func (d deck) print() { // d  because by convention we use one or two letters
	for i, card := range d {
		fmt.Println(i, card) //loop style with index
	}
}
// function to make hand deal cards 
func deal(d deck, handSize int) (deck, deck) { // (deck, deck) wanted to say it will return to value of deck type 
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string { // to transform our deck in string 
	return strings.Join([]string(d),",") // = we want to transform deck in one string and we choose for separator ","
	
}

func (d deck) saveToFile(filename string) error { // to save our data in file, deck because the date will be of type deck 
	return os.WriteFile(filename, []byte(d.toString()), 0666) // we call package os + we ask to transform in byte inside we transform our deck in one string + 0666 all permission to the file
}


func newDeckFromFile(filename string) deck { // to receive deck for receiver, no receiver for this function it is logical but the output will be of deck type 
	bs, err := os.ReadFile(filename) // bs = byte slice , err = error 
	if err != nil { // to check if no error and if errors the following code can start nil = absence 
		// option #1 - log the error and return a call to newDeck()
		// option #2 - Log the error and entirely quit the program ==> in this case we choose this one 
		fmt.Println("Error: ",err) // print errors in friendly format
		os.Exit(1) // 1 to say something was wrong 
	}

	s:= strings.Split(string(bs), ",") // for the "," in our case 
	return 	deck(s) // because we initiated deck like a slice of string 

}

// function to shuffle cards 
func (d deck) shuffle() { // no arguments are needed 
	
	source := rand.NewSource(time.Now().UnixNano()) // we created a new source for the seed of random
	r:=rand.New(source) // we stock the new seed in r 

	for i := range d {
		newPosition := r.Intn(len(d)-1) // to give a random index order in function of len of deck , we use the package math/rand
		
		d[i], d[newPosition] = d[newPosition], d[i] // basic line to shuffle 
	}

}