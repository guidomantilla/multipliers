package core

import (
	"fmt"
)

var multipliers = []int{3, 5}

func GetMultiplierType(number int) string {
	xxx := 0
	result := fmt.Sprintf("%d", number)
	for index, m := range multipliers {
		if number%m == 0 {
			xxx += index + 1
			result = fmt.Sprintf("Type %d", xxx)
		}
	}
	return result
}
