package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	input   string
	stdin   bool
	version bool
	args    []string
)

var (
	errTemplate = errors.New("template is required")
)

type options struct {
	data string
	tmpl []string
}

func init() {
	flag.Usage = usage
	flag.StringVar(&input, "d", "", "json data file")
	flag.BoolVar(&stdin, "i", false, "read data from stdin")
	flag.BoolVar(&version, "v", false, "print version info")
	flag.Parse()
	args = flag.Args()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] template [templates...]\n", NAME)
	flag.PrintDefaults()
}

func getOptions() (*options, error) {
	ops := new(options)
	if version {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if len(args) < 1 {
		return nil, errTemplate
	}
	ops.tmpl = args
	if stdin {
		ops.data = STDIN
	} else {
		ops.data = input
	}
	return ops, nil
}
