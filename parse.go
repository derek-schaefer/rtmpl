package main

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"text/template"
)

const (
	STDIN = "stdin"
)

type object map[string]interface{}

func parseTmpl(src []string) (*template.Template, error) {
	var (
		tmpl *template.Template
		err  error
	)
	tmpl = template.New(path.Base(src[0]))
	tmpl = tmpl.Funcs(helpers())
	_, err = tmpl.ParseFiles(src...)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func parseData(src string) (object, error) {
	var (
		read io.Reader
		data object
		err  error
	)
	if src == STDIN {
		read = os.Stdin
	} else if src != "" {
		var f *os.File
		f, err = os.Open(src)
		if f != nil {
			defer f.Close()
		}
		read = f
	}
	if read != nil && err == nil {
		err = json.NewDecoder(read).Decode(&data)
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
