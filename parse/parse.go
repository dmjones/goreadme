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

package parse

import (
	"bytes"
	"go/ast"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

func readExtraMarkdown(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func getPackageDocs(packageDir string) (docs, name, importPath string, err error) {
	var buildPkg *build.Package
	buildPkg, err = build.ImportDir(packageDir, build.ImportComment)
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

// ConvertDocs processes the package docs and converts to markdown. The
func ConvertDocs(packageDir string, c *Config) (string, error) {

	// Read additional markdown file, if present in config
	var customMarkdown string
	if c.CustomMarkdownFile != "" {
		var err error
		customMarkdown, err = readExtraMarkdown(c.CustomMarkdownFile)
		if err != nil {
			return "", err
		}
	}

	packageDocs, packageName, importPath, err := getPackageDocs(packageDir)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	ToMD(&buf, packageDocs, c.FencedCodeLanguage)
	markdownDocs := strings.TrimSpace(buf.String())

	tmpl, err := template.New("").Parse(pkgtemplate)
	if err != nil {
		return "", err
	}

	buf.Truncate(0)
	err = tmpl.Execute(&buf, struct {
		PackageName   string
		ImportPath    string
		PackageDocs   string
		ExtraMarkdown string
		Config        *Config
	}{
		PackageName:   packageName,
		ImportPath:    importPath,
		PackageDocs:   markdownDocs,
		ExtraMarkdown: customMarkdown,
		Config:        c,
	})

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
