package day3

import "testing"

// func TesttypeSwitch(t *testing.T) {
// 	a := typeSwitch(10)
// 	t.Log(a)
// }

func TestMap(t *testing.T) {
	numbers := []base{1, 2, 3, 4}
	strings := []base{"a", "b", "c"}
	newNumbers := Map(numbers, typeSwitch)
	newStrings := Map(strings, typeSwitch)
	t.Log(newNumbers)
	t.Log(newStrings)
}

func TestCat(t *testing.T) {
	Cat("log.txt", true)
}
