package main

import (
	"fmt"
	"time"
)

func main() {
	//https://go.dev/tour/basics/7

	//var card string = "Ace of Spades"
	/* card := newCard()

	fmt.Println(card) */

	defer fmt.Println("Esto es lo primero en declararse pero lo ultimo en salir por pantalla")

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)

	naked := nakedReturn("Esto te devuelve el mismo texto")
	fmt.Println(naked)

	var entero int = 1
	var enteroUnsigned uint = 1

	fmt.Println(entero)
	fmt.Println(enteroUnsigned)

	var enteroDeNuevo int = 12
	var ahoraEsFloat float64 = float64(enteroDeNuevo)
	fmt.Println(ahoraEsFloat)

	//enteroString := 42
	//cadena := string(enteroString)
	//fmt.Println("Para que se vea tienes que hacer fmt.Sprint(cadena) o usar strconv" + fmt.Sprint(cadena))

	//fmt.Println(cadena) // Imprimirá "4"

	//const constantes_bien string = "SOY UNA COSNTANTE"

	fmt.Println(needInt(1))
	fmt.Println(needFloat(12))

	//Principal()

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//while

	sumi := 1
	for sumi < 1000 {
		sumi += sumi
	}
	fmt.Println(sumi)

	/* for {
		fmt.Println("Bucle infinitoo")
	} */

	var condi int = 1

	if condi == 2 {
		fmt.Println("waos es 1")
	} else {
		fmt.Println("waos pos no es 1")
	}

	var dia = time.Now().Weekday()

	switch time.Sunday {
	case dia + 0:
		fmt.Println("Es hoy")
	case dia + 1:
		fmt.Println("Es mañana")
	default:
		fmt.Println("Todavia queda")
	}

}

func newCard() string {
	return "Five of Diamonds"
}

func nakedReturn(parameter string) (retorno string) {
	retorno = parameter
	return
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
