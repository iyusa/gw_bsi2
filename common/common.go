package common

import (
	"fmt"
	"log"
	"runtime"
)

// Version for AppServer
const Version = "0.3"
const showStackTrace = true

// Change log

// LogIfError print to stdErr if error occured
func LogIfError(err error) {
	if err != nil {
		log.Printf("Got Error: %v\n", err)
	}
}

// WarnIfError print to stdErr if error occured
func WarnIfError(err error) {
	if err != nil {
		log.Printf("WARNING: %v\n", err)
	}
}

// Ensure make sure err is not null, panic otherwise
func Ensure(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Catch error handler
//
// cara pakai: defer func() { err = common.Catch() }()
func Catch() (err error) {
	r := recover()
	if r != nil {
		err = r.(error)

		if showStackTrace {
			pc := make([]uintptr, 15)
			n := runtime.Callers(4, pc) // isi dengan depth stacktrace
			frames := runtime.CallersFrames(pc[:n])
			frame, _ := frames.Next()
			method := fmt.Sprintf("%s:%d", frame.Function, frame.Line)

			fmt.Printf("Ada Error @(%s) -> '%v'\n", method, err)
		}
	}
	return
}

// Validate on err
func Validate(err error, title string) {
	if err != nil {
		log.Printf("%s: %v\n", title, err)
		panic(err)
	}
}

// Assert condition to true, if false then panic
func Assert(condition bool, message string) {
	if !condition {
		err := fmt.Errorf(message)
		panic(err)
	}
}

// Abort force panic
func Abort(message string) {
	panic(fmt.Errorf(message))
}
