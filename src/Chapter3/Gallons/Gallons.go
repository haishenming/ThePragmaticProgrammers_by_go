package Gallons

func Gallons(piece, square int) int {

	gallonsNumber := square / piece
	hasGallons := square % piece
	if hasGallons == 0{
		return gallonsNumber
	} else {
		return gallonsNumber + 1
	}
}
