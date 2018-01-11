package Pizza

func DividedPizzas(pizzas, people, pieces int) (peoplePieces int, leftover int)  {

	totalPieces := pizzas * pieces
	peoplePieces = totalPieces / people
	leftover = totalPieces % people

	return
}
