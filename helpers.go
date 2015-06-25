package main

import (
	"os"
	"strings"
	"text/template"
)

func helpers() template.FuncMap {
	self := make(template.FuncMap)
	self["env"] = envHelper
	self["title"] = titleHelper
	self["upcase"] = upcaseHelper
	self["downcase"] = downcaseHelper
	return self
}

func envHelper(s string) string {
	return os.Getenv(s)
}

func titleHelper(s string) string {
	return strings.Title(s)
}

func upcaseHelper(s string) string {
	return strings.ToUpper(s)
}

func downcaseHelper(s string) string {
	return strings.ToLower(s)
}
