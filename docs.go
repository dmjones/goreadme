//   Copyright 2018 Duncan Jones
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

/*
goreadme converts Go package documentation into Github-friendly markdown,
designed to be used as a README.md file. There is support for adding badges
for documentation and build status.

Installation

TODO (need to configure goreleaser)

Usage

Run goreadme in your package directory and direct the output to "README.md":

  goreadme > README.md

The behaviour of the tool can be adjusted by including a `.goreadme.toml` file
in your package directory.

Some badges are supported directly by the tool (PRs welcome for others). For
those badges, just set the appropriate flags in the config file:

  # Shows a godoc badge for your package (https://godoc.org)
  showGodocBadge = true

  # Shows a Go report card for your package (https://goreportcard.com)
  showGoReportBadge = false

For others, specify the markdown in the config file:

  customMarkdownBadges = [
    "[![Coverage Status](https://link/to/badge.svg)](https://link/to/data)",
    "[![Build Status](https://link/to/badge.svg)](https://link/to/data)",
  ]

Additional markdown can be specified, which will be appended to the output from
the tool. To do this, give the name of a file containing the markdown to include.

  customMarkdownFile = ""

Acknowledgments

This tool leans heavily on the godoc->markdown work by Dave Cheney.
*/
package main
