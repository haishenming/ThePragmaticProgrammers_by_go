package Retire

import "time"

func OperationRetire(age int, retire int) (int, int, int)  {
	currentYear := time.Now().Year()
	hasYear := retire - age
	retireYear := currentYear + hasYear

	return hasYear, retireYear, currentYear

}
