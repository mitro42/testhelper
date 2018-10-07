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

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\texp: %#v\n\tgot: %#v\033[39m\n", filepath.Base(file), line, exp, act)
		tb.Fail()
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
