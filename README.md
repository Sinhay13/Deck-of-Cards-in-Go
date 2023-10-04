Deck of Cards in Go
This is a simple Go project that implements a deck of cards. It provides functionality to create, shuffle, and deal cards within a deck. It can also save and load a deck from a file.

Features
Create a deck of cards.
Shuffle the deck to randomize the card order.
Deal a specified number of cards from the deck.
Convert the deck to a string and save it to a file.
Load a deck from a file and restore its state.
Getting Started
Clone this repository to your local machine:

shell
Copy code
git clone https://github.com/sinhay13/deck-of-cards-go.git
Build and run the project:

shell
Copy code
cd deck-of-cards-go
go run main.go deck.go
Explore the project and use the provided functions to manipulate the deck of cards.

Usage
Here's an example of how to use the deck of cards:

go
Copy code
package main

import (
    "fmt"
)

func main() {
    // Create a new deck
    cards := newDeck()

    // Shuffle the deck
    cards.shuffle()

    // Deal some cards
    hand, remainingDeck := deal(cards, 5)

    // Print the hand and remaining deck
    fmt.Println("Hand:")
    hand.print()
    fmt.Println("\nRemaining Deck:")
    remainingDeck.print()
}
Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request if you have any improvements or bug fixes to suggest.

License
This project is licensed under the MIT License - see the LICENSE file for details.

