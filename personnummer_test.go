package personnummer

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestPersonnummerWithControlDigit(t *testing.T) {
	assert.Equal(t, true, Test(6403273813))
	assert.Equal(t, true, Test("510818-9167"))
	assert.Equal(t, true, Test("19900101-0017"))
	assert.Equal(t, true, Test("19130401+2931"))
	assert.Equal(t, true, Test("196408233234"))
}

func TestPersonnummerWithoutControlDigit(t *testing.T) {
	assert.Equal(t, false, Test(640327381))
	assert.Equal(t, false, Test("510818-916"))
	assert.Equal(t, false, Test("19900101-001"))
	assert.Equal(t, false, Test("100101+001"))
}

func TestPersonnummerWithWrongPersonnummerOrTypes(t *testing.T) {
	assert.Equal(t, false, Test([]string{}))
	assert.Equal(t, false, Test([]int{}))
	assert.Equal(t, false, Test(true))
	assert.Equal(t, false, Test(false))
	assert.Equal(t, false, Test(1122334455))
	assert.Equal(t, false, Test("112233-4455"))
	assert.Equal(t, false, Test("19112233-4455"))
	assert.Equal(t, false, Test("9999999999"))
	assert.Equal(t, false, Test("199999999999"))
	assert.Equal(t, false, Test("9913131315"))
	assert.Equal(t, false, Test("9911311232"))
	assert.Equal(t, false, Test("9902291237"))
	assert.Equal(t, false, Test("19990919_3766"))
	assert.Equal(t, false, Test("990919_3766"))
	assert.Equal(t, false, Test("199909193776"))
	assert.Equal(t, false, Test("Just a string"))
	assert.Equal(t, false, Test("990919+3776"))
	assert.Equal(t, false, Test("990919-3776"))
	assert.Equal(t, false, Test("9909193776"))
}

func TestCoOrdinationNumbers(t *testing.T) {
	assert.Equal(t, true, Test("701063-2391"))
	assert.Equal(t, true, Test("640883-3231"))
}

func TestWrongCoOrdinationNumbers(t *testing.T) {
	assert.Equal(t, false, Test("900161-0017"))
	assert.Equal(t, false, Test("640893-3231"))
}
