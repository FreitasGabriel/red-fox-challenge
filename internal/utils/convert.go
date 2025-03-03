package utils

import (
	"strconv"

	"github.com/FreitasGabriel/red-fox-challenge/config/logger"
)

func ConvertStringToInt(number string) int {
	if number == "" {
		return 0
	}
	convertedNumber, err := strconv.Atoi(number)
	if err != nil {
		logger.Error("error to convert number", err)
	}
	return convertedNumber
}
