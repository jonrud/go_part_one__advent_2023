package utils

import "strconv"

func IsNumber(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
