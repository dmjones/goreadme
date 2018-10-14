# goreadme
[![GoDoc](https://godoc.org/github.com/dmjones/goreadme?status.svg)](https://godoc.org/github.com/dmjones/goreadme)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmjones/goreadme)](https://goreportcard.com/report/github.com/dmjones/goreadme)
[![Build Status](https://travis-ci.com/dmjones/goreadme.svg?branch=master)](https://travis-ci.com/dmjones/goreadme)

goreadme converts Go package documentation into Github-friendly markdown,
designed to be used as a README.md file. There is support for adding badges
for documentation and build status.

### Installation
TODO (need to configure goreleaser)

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