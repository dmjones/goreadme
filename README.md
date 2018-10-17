# goreadme
[![GoDoc](https://godoc.org/github.com/dmjones/goreadme?status.svg)](https://godoc.org/github.com/dmjones/goreadme)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmjones/goreadme)](https://goreportcard.com/report/github.com/dmjones/goreadme)
[![Build Status](https://travis-ci.com/dmjones/goreadme.svg?branch=master)](https://travis-ci.com/dmjones/goreadme)

goreadme converts Go package documentation into a README.md file. This
avoids duplicating effort when writing docs and generally results in more
detailed Go package documentation. Win, win!

The Go documentation is parsed and converted into markdown. Build status badges
or documentation links can be added automatically and additional markdown can
be appended to the end of the file if needed.

The README.md for this project is generated using this tool. See `docs.go` for
the source material and `.goreadme.toml` for the configuration. More details on
these below.

### Installation
It is recommended to download a binary from
<a href="https://github.com/dmjones/goreadme/releases">https://github.com/dmjones/goreadme/releases</a>.

If you'd like to build from the latest source you must first
install dep (<a href="https://golang.github.io/dep/docs/installation.html">https://golang.github.io/dep/docs/installation.html</a>). Then
run these commands:


```
go get -d github.com/dmjones/goreadme
cd "$GOPATH/src/github.com/dmjones/goreadme"
dep ensure --vendor-only
go install .
```

### Usage
Run goreadme in your package directory and direct the output to "README.md":


```
goreadme README.md
```

Omitting the file name will print to stdout.

### Configuration
The behaviour of the tool can be adjusted by including a `.goreadme.toml` file
in your package directory.

Some badges are supported directly by the tool (PRs welcome for others). For
those badges, just set the appropriate flags in the config file:


```
# Shows a godoc badge for your package
showGodocBadge = true

# Shows a Go report card for your package
showGoReportBadge = true
```

For others, specify the markdown in the config file:


```
customMarkdownBadges = [
  "[![Coverage Status](...)](...)",
  "[![Build Status](...)](...)",
]
```

Additional markdown can be specified, which will be appended to the output from
the tool. To do this, give the name of a file containing the markdown to include.


```
customMarkdownFile = "extraInfo.md"
```

### Acknowledgments
This tool is based on the <a href="https://github.com/davecheney/godoc2md">https://github.com/davecheney/godoc2md</a> project by Dave
Cheney.

<sub>*generated with [goreadme](https://github.com/dmjones/goreadme)*</sub>