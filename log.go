package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kovetskiy/lorg"
	"github.com/kovetskiy/spinner-go"
)

func getLogger() *lorg.Log {
	logger := lorg.NewLog()
	logger.SetFormat(lorg.NewFormat("${level:[%s]:right:short} %s"))

	return logger
}

// use fatalf or fatalln
func fatal(format string, values ...interface{}) {
	if spinner.IsActive() {
		spinner.Stop()
	}

	_, file, ln, ok := runtime.Caller(2)
	if ok {
		format = fmt.Sprintf("%s#%d: %v", file, ln, format)
	}
	fmt.Fprintf(os.Stderr, format+"\n", values...)
	os.Exit(1)
}

func fatalf(format string, values ...interface{}) {
	fatal(format, values...)
}

func fatalln(value interface{}) {
	fatal("%s", value)
}

// use debugf or debugln
func debug(format string, values ...interface{}) {
	_, file, ln, ok := runtime.Caller(2)
	if ok {
		format = fmt.Sprintf("%s#%d: %v", file, ln, format)
	}
	logger.Debugf(format, values...)
}

func debugf(format string, values ...interface{}) {
	debug(format, values...)
}

func debugln(value interface{}) {
	debug("%s", value)
}

// use tracef or traceln
func trace(format string, values ...interface{}) {
	_, file, ln, ok := runtime.Caller(2)
	if ok {
		format = fmt.Sprintf("%s#%d: %v", file, ln, format)
	}
	logger.Tracef(format, values...)
}

func tracef(format string, values ...interface{}) {
	trace(format, values...)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func traceln(value interface{}) {
	trace("%s", value)
}
