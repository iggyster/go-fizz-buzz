package fizzbuzz

import (
	"strconv"
)

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

type RoundsNumber uint

func (number RoundsNumber) String() string {
	return strconv.Itoa(int(number))
}

type RoundsHistory []string

func (history RoundsHistory) add(index int, input string) {
	history[index] = input
}

func Play(len int) RoundsHistory {
	history := make(RoundsHistory, len)
	for index := range history {
		num := RoundsNumber(1 + index)
		switch {
		case isFizzBuzz(num):
			history.add(index, Fizz+Buzz)
		case isFizz(num):
			history.add(index, Fizz)
		case isBuzz(num):
			history.add(index, Buzz)
		default:
			history.add(index, num.String())
		}
	}
	return history
}

func isFizzBuzz(num RoundsNumber) bool {
	return isFizz(num) && isBuzz(num)
}

func isFizz(num RoundsNumber) bool {
	return num%3 == 0
}

func isBuzz(num RoundsNumber) bool {
	return num%5 == 0
}
