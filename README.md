# main
[![GoDoc](https://godoc.org/github.com/dmjones/goreadme?status.svg)](https://godoc.org/github.com/dmjones/goreadme)


goreadme converts Go package documentation into Github-friendly markdown,
designed to be used as a README.md file. There is support for adding badges
for documentation and build status.

### Installation
TODO (need to configure goreleaser)

### Usage
Run goreadme in your package directory and direct the output to "README.md":


```
goreadme > README.md
```

The behaviour of the tool can be adjusted by including a `.goreadme.toml` file
in your package directory.

Some badges are supported directly by the tool (PRs welcome for others). For
those badges, just set the appropriate flags in the config file:


```
# Shows a godoc badge for your package (<a href="https://godoc.org">https://godoc.org</a>)
showGodocBadge = true

# Shows a Go report card for your package (<a href="https://goreportcard.com">https://goreportcard.com</a>)
showGoReportBadge = false
```

For others, specify the markdown in the config file:


```
customMarkdownBadges = [
  "[![Coverage Status](<a href="https://link/to/badge.svg">https://link/to/badge.svg</a>)](<a href="https://link/to/data">https://link/to/data</a>)",
  "[![Build Status](<a href="https://link/to/badge.svg">https://link/to/badge.svg</a>)](<a href="https://link/to/data">https://link/to/data</a>)",
]
```

Additional markdown can be specified, which will be appended to the output from
the tool. To do this, give the name of a file containing the markdown to include.


```
customMarkdownFile = ""
```

### Acknowledgments
This tool leans heavily on the godoc->markdown work by Dave Cheney.




<sub>*generated with [go2readme](https://github.com/dmjones/go2readme)*</sub>
