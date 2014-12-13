package personnummer

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestPersonnummerWithControlDigit(t *testing.T) {
	assert.Equal(t, true, Valid(6403273813))
	assert.Equal(t, true, Valid("510818-9167"))
	assert.Equal(t, true, Valid("19900101-0017"))
	assert.Equal(t, true, Valid("19130401+2931"))
	assert.Equal(t, true, Valid("196408233234"))
}

func TestPersonnummerWithoutControlDigit(t *testing.T) {
	assert.Equal(t, false, Valid(640327381))
	assert.Equal(t, false, Valid("510818-916"))
	assert.Equal(t, false, Valid("19900101-001"))
	assert.Equal(t, false, Valid("100101+001"))
}

func TestPersonnummerWithWrongPersonnummerOrTypes(t *testing.T) {
	assert.Equal(t, false, Valid([]string{}))
	assert.Equal(t, false, Valid([]int{}))
	assert.Equal(t, false, Valid(true))
	assert.Equal(t, false, Valid(false))
	assert.Equal(t, false, Valid(1122334455))
	assert.Equal(t, false, Valid("112233-4455"))
	assert.Equal(t, false, Valid("19112233-4455"))
	assert.Equal(t, false, Valid("9999999999"))
	assert.Equal(t, false, Valid("199999999999"))
	assert.Equal(t, false, Valid("9913131315"))
	assert.Equal(t, false, Valid("9911311232"))
	assert.Equal(t, false, Valid("9902291237"))
	assert.Equal(t, false, Valid("19990919_3766"))
	assert.Equal(t, false, Valid("990919_3766"))
	assert.Equal(t, false, Valid("199909193776"))
	assert.Equal(t, false, Valid("Just a string"))
	assert.Equal(t, false, Valid("990919+3776"))
	assert.Equal(t, false, Valid("990919-3776"))
	assert.Equal(t, false, Valid("9909193776"))
}

func TestCoOrdinationNumbers(t *testing.T) {
	assert.Equal(t, true, Valid("701063-2391"))
	assert.Equal(t, true, Valid("640883-3231"))
}

func TestWrongCoOrdinationNumbers(t *testing.T) {
	assert.Equal(t, false, Valid("900161-0017"))
	assert.Equal(t, false, Valid("640893-3231"))
}
