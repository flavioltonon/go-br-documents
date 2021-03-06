package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// PIS struct
type pis struct {
	number pisNumber
	valid  bool
}

func (p pis) Number(mask bool) string {
	if p.valid && mask {
		return string(p.number[:3]) + "." + string(p.number[3:8]) + "." + string(p.number[8:10]) + "-" + string(p.number[10:])
	}
	return string(p.number)
}

func ParsePIS(number string) (pis, error) {
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")
	if len(number) != 11 && len(number) != 13 {
		return pis{}, ErrIncorrectLenghtPISNumber
	}

	pisNumber := pisNumber(number)

	if pisNumber.isFalsePositive() {
		return pis{}, ErrInvalidPISNumber
	}

	if !pisNumber.hasValidDigit() {
		return pis{}, ErrInvalidPISNumber
	}

	return pis{
		number: pisNumber,
		valid:  true,
	}, nil
}

func RandomPISNumber(mask bool) string {
	var (
		source      = rand.NewSource(time.Now().UnixNano())
		multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	r := rand.New(source)
	pisNumber := int(r.Int63n(8999999999) + 1000000000)
	pisString := strconv.Itoa(pisNumber)

	for i := 0; i < 10; i++ {
		number, _ := strconv.Atoi(string(pisString[i]))
		sum += number * multipliers[i]
	}
	digit := 11 - sum%11
	if digit >= 10 {
		digit = 0
	}

	if mask {
		return pisString[:3] + "." + pisString[3:8] + "." + pisString[8:] + "-" + strconv.Itoa(digit)
	}
	return pisString + strconv.Itoa(digit)
}

type pisNumber string

func (p pisNumber) isFalsePositive() bool {
	if string(p) == "00000000000" {
		return true
	}
	return false
}

func (p pisNumber) hasValidDigit() bool {
	var (
		multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	for i := 0; i < 10; i++ {
		pisDigit, _ := strconv.Atoi(string(p[i]))
		sum += pisDigit * multipliers[i]
	}

	return string(p[10]) == strconv.Itoa((11-sum%11)%10) || string(p[10]) == strconv.Itoa((11-sum%11)%11)
}
