package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type SpeSkillTest struct {
}

func TestNarsisticNumber(t *testing.T) {
	test := SpeSkillTest{}
	assert.True(t, test.NarsisticNumber(1634))
}

func (spe SpeSkillTest) NarsisticNumber(number int) bool {
	numberStr := strconv.Itoa(number)
	total := 0
	numberLength := len(numberStr)
	for i := 0; i < numberLength; i++ {
		myInt, err := strconv.Atoi(numberStr[i : i+1])
		if err != nil {
			total += myInt ^ numberLength
		}
	}
	return number == total
}
