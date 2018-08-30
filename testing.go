package testing

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.Fail()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n", filepath.Base(file), line, err.Error())
		tb.Fail()
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

func expectPanic(tb testing.TB, expectedMessage string) {
	r := recover()
	_, file, line, _ := runtime.Caller(1)
	if r == nil {
		fmt.Printf("\033[31m%s:%d:\nExpected panic with error message: \n\t%v\033[39m\n",
			filepath.Base(file), line, expectedMessage)
		tb.Fail()
	} else if r != expectedMessage {
		fmt.Printf("\033[31m%s:%d:\nExpected panic with error message: \n\texp: %v\n\tgot: %v\033[39m\n",
			filepath.Base(file), line, expectedMessage, r)
		tb.Fail()
	}
}
