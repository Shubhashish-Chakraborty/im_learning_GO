package helper

import (
	"math/rand"
)

func RandomNumber() int {
	ranNumber := rand.Intn(1000); // random number from 0 to 999
	return ranNumber;
}