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
for documentation and build status. The README.md for this project is generated
using this tool.

Installation

It is recommended to download a binary from
https://github.com/dmjones/goreadme/releases.

If you'd like to build from the latest source you must first
install dep (https://golang.github.io/dep/docs/installation.html). Then
run these commands:

  go get -d github.com/dmjones/goreadme
  cd "$GOPATH/src/github.com/dmjones/goreadme"
  dep ensure --vendor-only
  go install .

Usage

Run goreadme in your package directory and direct the output to "README.md":

  goreadme README.md

Omitting the file name will print to stdout.

Configuration

The behaviour of the tool can be adjusted by including a `.goreadme.toml` file
in your package directory.

Some badges are supported directly by the tool (PRs welcome for others). For
those badges, just set the appropriate flags in the config file:

  # Shows a godoc badge for your package
  showGodocBadge = true

  # Shows a Go report card for your package
  showGoReportBadge = true

For others, specify the markdown in the config file:

  customMarkdownBadges = [
    "[![Coverage Status](...)](...)",
    "[![Build Status](...)](...)",
  ]

Additional markdown can be specified, which will be appended to the output from
the tool. To do this, give the name of a file containing the markdown to include.

  customMarkdownFile = "extraInfo.md"

Acknowledgments

This tool is based on the https://github.com/davecheney/godoc2md project by Dave
Cheney.
*/
package main

//go:generate goreadme README.md
