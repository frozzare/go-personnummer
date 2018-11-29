package personnummer

import "fmt"

const (
	lengthWithoutCentury = 10
	lengthWithCentury    = 12
)

// ValidateStrings validate Swedish social security numbers.
func ValidString(in string) bool {
	cleanNumber := make([]byte, 0, len(in))
	for _, c := range in {
		if c == '+' {
			continue
		}
		if c == '-' {
			continue
		}

		if c > '9' {
			return false
		}
		if c < '0' {
			return false
		}

		cleanNumber = append(cleanNumber, byte(c))
	}

	switch len(cleanNumber) {
	case lengthWithCentury:
		if !luhn(cleanNumber[2:]) {
			return false
		}

		dateBytes := append(cleanNumber[:6], getCoOrdinationDay(cleanNumber[6:8])...)
		return validateTime(dateBytes)
	case lengthWithoutCentury:
		if !luhn(cleanNumber) {
			return false
		}

		dateBytes := append(cleanNumber[:4], getCoOrdinationDay(cleanNumber[4:6])...)
		return validateTime(dateBytes)
	default:
		return false
	}
}

var monthDays = map[int]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

// input time without centry.
func validateTime(time []byte) bool {
	length := len(time)

	date := charsToDigit(time[length-2 : length])
	month := charsToDigit(time[length-4 : length-2])

	if month != 2 {
		days, ok := monthDays[month]
		if !ok {
			return false
		}
		return date <= days
	}

	year := charsToDigit(time[:length-4])

	leapYear := year%4 == 0 && year%100 != 0 || year%400 == 0

	if leapYear {
		return date <= 29
	}
	return date <= 28
}

// Valid will validate Swedish social security numbers.
func Valid(i interface{}) bool {
	switch v := i.(type) {
	case int, int32, int64, uint, uint32, uint64:
		return ValidString(fmt.Sprint(v))
	case string:
		return ValidString(v)
	default:
		return false
	}
}

var rule3 = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// luhn will test if the given string is a valid luhn string.
func luhn(s []byte) bool {
	odd := len(s) & 1

	var sum int

	for i, c := range s {
		if i&1 == odd {
			sum += rule3[c-'0']
		} else {
			sum += int(c - '0')
		}
	}

	return sum%10 == 0
}

// getCoOrdinationDay will return co-ordination day.
func getCoOrdinationDay(day []byte) []byte {
	d := charsToDigit(day)
	if d < 60 {
		return day
	}

	d -= 60

	if d < 10 {
		return []byte{'0', byte(d) + '0'}
	}

	return []byte{
		byte(d)/10 + '0',
		byte(d)%10 + '0',
	}
}

// charsToDigit converts char bytes to a digit
// example: ['1', '1'] => 11
func charsToDigit(chars []byte) int {
	l := len(chars)
	r := 0
	for i, c := range chars {
		p := int((c - '0'))
		for j := 0; j < l-i-1; j++ {
			p *= 10
		}
		r += p
	}
	return r
}
