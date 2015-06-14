package main

import (
	"os"
)

var (
	NAME    = "rtmpl"
	VERSION = "HEAD"
)

func main() {
	ops, err := getOptions()
	exitError(err)
	tmpl, err := parseTmpl(ops.tmpl)
	exitError(err)
	data, err := parseData(ops.data)
	exitError(err)
	err = tmpl.Execute(os.Stdout, data)
	exitError(err)
}
