package main

import "fmt"

func main() {
	//https://go.dev/tour/basics/7

	//var card string = "Ace of Spades"
	/* card := newCard()

	fmt.Println(card) */

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)

	naked := nakedReturn("Esto te devuelve el mismo texto")
	fmt.Println(naked)

}

func newCard() string {
	return "Five of Diamonds"
}

func nakedReturn(parameter string) (retorno string) {
	retorno = parameter
	return
}
