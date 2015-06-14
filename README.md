# rtmpl

A command-line template rendering utility.

## Usage

Create a template.

```bash
$ cat example.tmpl
{{ env "NAME" }}
{{ range .numbers }}
  {{ . }}
{{ end }}
```

Create some JSON data.

```bash
$ cat example.json
{
  "numbers": [1, 2, 3]
}
```

Render the template.

```bash
$ export NAME=test
$ rtmpl -d example.json example.tmpl
test
  1
  2
  3
```

Data can also be read from stdin.

```bash
$ export NAME=test
$ cat example.json | rtmpl -i example.tmpl
test
  1
  2
  3
```

Multiple templates can be included:

```bash
$ cat example1.tmpl
{{ env "NAME" }}
{{ template "example2.tmpl" . }}
$ cat example2.tmpl
{{ range .numbers }}
  {{ . }}
{{ end }}
$ export NAME=test
$ rtmpl -d example.json example1.tmpl example2.tmpl
test
  1
  2
  3
```

## Installation

Binary:

Check the "releases" page.

Go:

`$ go get github.com/derek-schaefer/rtmpl`

Docker:

`$ docker pull derekschaefer/rtmpl`

## Template Reference

The template syntax is provided by Go's [text.Template](http://golang.org/pkg/text/template/) package.

Additional helpers:

+ `env(string) string`: returns an environment variable by name

## License

MIT
