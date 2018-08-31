package personnummer

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestPersonnummerWithControlDigit(t *testing.T) {
	assert.True(t, Valid(6403273813))
	assert.True(t, Valid("510818-9167"))
	assert.True(t, Valid("19900101-0017"))
	assert.True(t, Valid("19130401+2931"))
	assert.True(t, Valid("196408233234"))
	assert.True(t, Valid("0001010107"))
	assert.True(t, Valid("000101-0107"))
	assert.True(t, Valid("200002296127"))
}

func TestPersonnummerWithoutControlDigit(t *testing.T) {
	assert.False(t, Valid(640327381))
	assert.False(t, Valid("510818-916"))
	assert.False(t, Valid("19900101-001"))
	assert.False(t, Valid("100101+001"))
}

func TestPersonnummerWithWrongPersonnummerOrTypes(t *testing.T) {
	assert.False(t, Valid([]string{}))
	assert.False(t, Valid([]int{}))
	assert.False(t, Valid(true))
	assert.False(t, Valid(false))
	assert.False(t, Valid(1122334455))
	assert.False(t, Valid("112233-4455"))
	assert.False(t, Valid("19112233-4455"))
	assert.False(t, Valid("9999999999"))
	assert.False(t, Valid("199999999999"))
	assert.False(t, Valid("9913131315"))
	assert.False(t, Valid("9911311232"))
	assert.False(t, Valid("9902291237"))
	assert.False(t, Valid("19990919_3766"))
	assert.False(t, Valid("990919_3766"))
	assert.False(t, Valid("199909193776"))
	assert.False(t, Valid("Just a string"))
	assert.False(t, Valid("990919+3776"))
	assert.False(t, Valid("990919-3776"))
	assert.False(t, Valid("9909193776"))
}

func TestCoOrdinationNumbers(t *testing.T) {
	assert.True(t, Valid("701063-2391"))
	assert.True(t, Valid("640883-3231"))
}

func TestWrongCoOrdinationNumbers(t *testing.T) {
	assert.False(t, Valid("900161-0017"))
	assert.False(t, Valid("640893-3231"))
}
