package main

import (
	"fmt"
	"os"
)

func exitError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
