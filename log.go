package main

import (
	"fmt"
	"os"

	"github.com/kovetskiy/lorg"
	"github.com/kovetskiy/spinner-go"
)

func getLogger() *lorg.Log {
	logger := lorg.NewLog()
	logger.SetFormat(lorg.NewFormat("${level:[%s]:left:true} %s"))

	return logger
}

func fatalf(format string, values ...interface{}) {
	if spinner.IsActive() {
		spinner.Stop()
	}

	fmt.Fprintf(os.Stderr, format+"\n", values...)
	os.Exit(1)
}

func fatalln(value interface{}) {
	fatalf("%s", value)
}

func debugf(format string, values ...interface{}) {
	logger.Debugf(format, values...)
}

func tracef(format string, values ...interface{}) {
	logger.Tracef(format, values...)
}

func debugln(value interface{}) {
	debugf("%s", value)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func traceln(value interface{}) {
	tracef("%s", value)
}
