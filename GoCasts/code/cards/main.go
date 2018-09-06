package main

func main() {
/*	cards := deck{"1", "2"}
	cards = append(cards, "3")
	cards.print()*/

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//hand, remainingCards := deal(cards, 5)
	//cards.saveToFile("myfile.txt")
}


