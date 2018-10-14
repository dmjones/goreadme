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

package main

import (
	"bytes"
	"errors"
	"go/ast"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/BurntSushi/toml"
)

const configFile = ".goreadme.toml"

type config struct {
	FencedCodeLanguage   string
	ShowGeneratedSuffix  bool
	ShowGodocBadge       bool
	ShowGoReportBadge    bool
	CustomMarkdownBadges []string
	CustomMarkdownFile   string
}

func defaultConfig() *config {
	return &config{
		FencedCodeLanguage:   "go",
		ShowGeneratedSuffix:  true,
		ShowGodocBadge:       true,
		ShowGoReportBadge:    false,
		CustomMarkdownBadges: nil,
		CustomMarkdownFile:   "",
	}
}

func main() {
	config, err := readConfig()
	logFatal(err, "Failed to read config")

	err = parse(config)
	logFatal(err, "Failed to parse package docs")
}

func readConfig() (*config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return defaultConfig(), nil
	}

	// Begin with default config. User can then overwrite what they want.
	c := defaultConfig()
	_, err := toml.DecodeFile(configFile, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func readExtraMarkdown(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func getPackageDocs() (docs, name, importPath string, err error) {
	var wd string
	wd, err = os.Getwd()
	if err != nil {
		return
	}

	var buildPkg *build.Package
	buildPkg, err = build.ImportDir(wd, build.ImportComment)
	if err != nil {
		return
	}

	fs := token.NewFileSet()

	filter := func(info os.FileInfo) bool {
		for _, name := range buildPkg.GoFiles {
			if name == info.Name() {
				return true
			}
		}
		return false
	}

	var pkgs map[string]*ast.Package
	pkgs, err = parser.ParseDir(fs, buildPkg.Dir, filter, parser.ParseComments)
	if err != nil {
		return
	}

	if len(pkgs) != 1 {
		err = errors.New("multiple packages found in directory")
		return
	}

	docPkg := doc.New(pkgs[buildPkg.Name], buildPkg.ImportPath, 0)

	docs = docPkg.Doc

	if buildPkg.Name == "main" {
		// In 99% of cases, this is a tool. We want the name of the containing
		// directory instead.
		name = path.Base(buildPkg.ImportPath)
	} else {
		name = buildPkg.Name
	}
	importPath = buildPkg.ImportPath
	return
}

func parse(c *config) error {

	// Read additional markdown file, if present in config
	var customMarkdown string
	if c.CustomMarkdownFile != "" {
		var err error
		customMarkdown, err = readExtraMarkdown(c.CustomMarkdownFile)
		if err != nil {
			return err
		}
	}

	packageDocs, packageName, importPath, err := getPackageDocs()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	ToMD(&buf, packageDocs, c.FencedCodeLanguage)
	markdownDocs := buf.String()

	tmpl, err := template.New("").Parse(pkgtemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, struct {
		PackageName   string
		ImportPath    string
		PackageDocs   string
		ExtraMarkdown string
		Config        *config
	}{
		PackageName:   packageName,
		ImportPath:    importPath,
		PackageDocs:   markdownDocs,
		ExtraMarkdown: customMarkdown,
		Config:        c,
	})

	if err != nil {
		return err
	}

	return nil
}

func logFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg + ": " + err.Error())
	}
}
