package main

import (
	"os"
	"text/template"
)

func helpers() template.FuncMap {
	self := make(template.FuncMap)
	self["env"] = envHelper
	return self
}

func envHelper(key string) string {
	return os.Getenv(key)
}
