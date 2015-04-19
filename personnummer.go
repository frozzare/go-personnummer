package personnummer

import (
	"math"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// The Luhn algorithm.
func luhn(str string) int {
	sum := 0

	for i, r := range str {
		c := string(r)
		v, _ := strconv.Atoi(c)
		v *= 2 - (i % 2)
		if v > 9 {
			v -= 9
		}
		sum += v
	}

	return int(math.Ceil(float64(sum)/10)*10 - float64(sum))
}

// Test date if lunh is true
func testDate(century string, year string, month string, day string) bool {
	t, err := time.Parse("01/02/2006", month+"/"+day+"/"+century+year)

	if err != nil {
		return false
	}

	y, _ := strconv.Atoi(century + year)
	m, _ := strconv.Atoi(month)
	d, _ := strconv.Atoi(day)

	if y > time.Now().Year() {
		return false
	}

	return !(t.Year() != y || int(t.Month()) != m || t.Day() != d)
}

// Get co-ordination number day
func getCoOrdinationDay(day string) string {
	d, _ := strconv.Atoi(day)
	d -= 60
	day = strconv.Itoa(d)

	if d < 10 {
		day = "0" + day
	}

	return day
}

// Validate Swedish personal identity numbers.
func Valid(str interface{}) bool {
	if reflect.TypeOf(str).Kind() != reflect.Int && reflect.TypeOf(str).Kind() != reflect.String {
		return false
	}

	pr := ""

	if reflect.TypeOf(str).Kind() == reflect.Int {
		pr = strconv.Itoa(str.(int))
	} else {
		pr = str.(string)
	}

	re, _ := regexp.Compile(`^(\d{2}){0,1}(\d{2})(\d{2})(\d{2})([\-|\+]{0,1})?(\d{3})(\d{0,1})$`)
	match := re.FindStringSubmatch(pr)

	if len(match) == 0 {
		return false
	}

	century := match[1]
	year := match[2]
	month := match[3]
	day := match[4]
	num := match[6]
	check := match[7]

	if len(century) == 0 {
		yearNow := time.Now().Year()
		years := [...]int{yearNow, yearNow - 100, yearNow - 150}

		for _, yi := range years {
			ys := strconv.Itoa(yi)

			if Valid(ys[:2] + pr) {
				return true
			}
		}

		return false
	}

	if len(year) == 4 {
		year = year[2:]
	}

	c, _ := strconv.Atoi(check)

	valid := luhn(year+month+day+num) == c && len(check) != 0

	if valid && testDate(century, year, month, day) {
		return valid
	}

	day = getCoOrdinationDay(day)

	return valid && testDate(century, year, month, day)
}
