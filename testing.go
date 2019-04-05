package testhelper

import (
	"reflect"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	tb.Helper()
	if !condition {
		tb.Errorf("\033[31m "+msg+"\033[39m\n", v...)
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Errorf("\033[31m unexpected error: %s\033[39m\n", err.Error())
	}
}

// Nok fails the test if err is nil, or the error message is different from the expected
func Nok(tb testing.TB, err error, expectedMessage string) {
	tb.Helper()
	if err == nil {
		tb.Errorf("\033[31m unexpected success, error shouldn't be nil")
	} else if err.Error() != expectedMessage {
		tb.Errorf("\033[31m unexpected error message: \n\texp: %#v\n\tgot: %#v\033[39m\n", expectedMessage, err.Error())
	}
}

// NokPrefix fails the test if err is nil, or the error message doesn't start with th required prefix
func NokPrefix(tb testing.TB, err error, expectedPrefix string) {
	tb.Helper()
	if err == nil {
		tb.Errorf("\033[31m unexpected success, error shouldn't be nil")
	} else if !strings.HasPrefix(err.Error(), expectedPrefix) {
		tb.Errorf("\033[31m unexpected error message: \n\texp: %#v\n\tgot: %#v\033[39m\n", expectedPrefix, err.Error())
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Errorf("\033[31m\n\texp: %#v (%T)\n\tgot: %#v (%T)\033[39m\n", exp, exp, act, act)
	}
}

// ExpectPanic fails if the test completes without a panic or with a panic but with incorrect error message
func ExpectPanic(tb testing.TB, expectedMessage string) {
	tb.Helper()
	r := recover()
	if r == nil {
		tb.Errorf("\033[31m\nExpected panic with error message: \n\t%v\033[39m\n", expectedMessage)
	} else if r != expectedMessage {
		tb.Errorf("\033[31m\nExpected panic with error message: \n\texp: %v\n\tgot: %v\033[39m\n", expectedMessage, r)
	}
}
