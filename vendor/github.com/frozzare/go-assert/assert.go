package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}

	if len(msgAndArgs) == 1 {
		return msgAndArgs[0].(string)
	}

	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}

	return ""
}

func equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	if fmt.Sprintf("%#v", expected) == fmt.Sprintf("%#v", actual) {
		return true
	}

	return false
}

func isnil(actual interface{}) bool {
	if actual == nil {
		return true
	}

	value := reflect.ValueOf(actual)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

var numericZeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

func empty(actual interface{}) bool {
	if isnil(actual) {
		return true
	}

	if actual == "" {
		return true
	}

	if actual == false {
		return true
	}

	for _, v := range numericZeros {
		if actual == v {
			return true
		}
	}

	actValue := reflect.ValueOf(actual)

	switch actValue.Kind() {
	case reflect.Map:
		fallthrough
	case reflect.Slice, reflect.Chan:
		return (actValue.Len() == 0)
	case reflect.Struct:
		switch actual.(type) {
		case time.Time:
			return actual.(time.Time).IsZero()
		}
	case reflect.Ptr:
		if actValue.IsNil() {
			return true
		}
		switch actual.(type) {
		case *time.Time:
			return actual.(*time.Time).IsZero()
		default:
			return false
		}
	}

	return false
}

// Fail logs a failed message.
func Fail(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	message := messageFromMsgAndArgs(msgAndArgs...)
	_, file, line, _ := runtime.Caller(2)

	if message > "" {
		message = fmt.Sprintf("(%s)", message)
	}

	t.Errorf("\r\033[39m%s:%d \033[31m%v == %v\033[39m %s\n",
		filepath.Base(file),
		line,
		expected,
		actual,
		message)
}

// Equal compare the expected value with the actual value and determine if the values are the same.
func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	success := equal(expected, actual)

	if !success {
		Fail(t, expected, actual, msgAndArgs...)
	}

	return success
}

// NotEqual compare the expected value with the actual value and determine if the values are not the same.
func NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	success := !equal(expected, actual)

	if !success {
		Fail(t, expected, actual, msgAndArgs...)
	}

	return success
}

// True checks if the given value is true or not.
func True(t *testing.T, actual bool, msgAndArgs ...interface{}) bool {
	success := actual == true

	if !success {
		Fail(t, true, actual, msgAndArgs...)
	}

	return success
}

// False checks if the given value is false or not.
func False(t *testing.T, actual bool, msgAndArgs ...interface{}) bool {
	success := actual == false

	if !success {
		Fail(t, false, actual, msgAndArgs...)
	}

	return success
}

// NotNil checks if the given value is not nil or not.
func NotNil(t *testing.T, actual interface{}, msgAndArgs ...interface{}) bool {
	success := !isnil(actual)

	if !success {
		Fail(t, nil, actual, msgAndArgs...)
	}

	return success
}

// Nil checks if the given value is nil or not.
func Nil(t *testing.T, actual interface{}, msgAndArgs ...interface{}) bool {
	success := isnil(actual)

	if !success {
		Fail(t, nil, actual, msgAndArgs...)
	}

	return success
}

// Empty checks if the given value is empty or not.
func Empty(t *testing.T, actual interface{}, msgAndArgs ...interface{}) bool {
	if empty(actual) {
		return true
	}

	Fail(t, "empty", actual, msgAndArgs...)

	return false
}

// NotEmpty checks if the given value is not empty or not.
func NotEmpty(t *testing.T, actual interface{}, msgAndArgs ...interface{}) bool {
	if !empty(actual) {
		return true
	}

	Fail(t, "not empty", actual, msgAndArgs...)

	return false
}
