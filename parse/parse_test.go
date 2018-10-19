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
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	runConfigTest(t, DefaultConfig(), "testdata/defaultConfig.md")
}

func TestNoSuffix(t *testing.T) {
	config := DefaultConfig()
	config.ShowGeneratedSuffix = false

	runConfigTest(t, config, "testdata/noSuffix.md")
}

func TestFencedCode(t *testing.T) {
	config := DefaultConfig()
	config.FencedCodeLanguage = "ruby"

	runConfigTest(t, config, "testdata/ruby.md")
}

func TestExtraMarkdown(t *testing.T) {
	config := DefaultConfig()
	config.CustomMarkdownFile = "testdata/extra.md"

	runConfigTest(t, config, "testdata/customMarkdown.md")
}

func TestCustomBadges(t *testing.T) {
	config := DefaultConfig()
	config.CustomMarkdownBadges = []string{"badge1", "badge2"}

	runConfigTest(t, config, "testdata/customBadges.md")
}

func TestCustomBadgesOnly(t *testing.T) {
	config := DefaultConfig()
	config.ShowGodocBadge = false
	config.CustomMarkdownBadges = []string{"badge1", "badge2"}

	runConfigTest(t, config, "testdata/customBadgesOnly.md")
}

func TestNoGodoc(t *testing.T) {
	config := DefaultConfig()
	config.ShowGodocBadge = false

	runConfigTest(t, config, "testdata/noGodoc.md")
}

func TestGoReport(t *testing.T) {
	config := DefaultConfig()
	config.ShowGoReportBadge = true

	runConfigTest(t, config, "testdata/goReport.md")
}

func runConfigTest(t *testing.T, config *Config, expectedFile string) {
	tmpDir := copyTestFileToGoPath(t, "foo/docs.go", "")
	defer os.RemoveAll(tmpDir)

	expectedBytes, err := ioutil.ReadFile(expectedFile)
	require.NoError(t, err)

	importPath := path.Base(tmpDir)
	expected := strings.Replace(string(expectedBytes), "%%IMPORT_PATH%%", importPath, -1)

	out, err := ConvertDocs(tmpDir, config)

	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestBadExtraMarkdownReturnsError(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	config := &Config{
		CustomMarkdownFile: "doesn't exist",
	}

	_, err = ConvertDocs(wd+"/foo", config)
	assert.Error(t, err)
}

func TestMainPackagesTreatedAsCommands(t *testing.T) {

	tmpDir := copyTestFileToGoPath(t, "bar/docs_main.go", "")
	defer os.RemoveAll(tmpDir)

	res, err := ConvertDocs(tmpDir, &Config{
		ShowGodocBadge:      false,
		ShowGeneratedSuffix: false,
	})
	require.NoError(t, err)

	// Since this test file was in the main package, the title of the page should
	// be the directory name.
	expected := fmt.Sprintf("# %s\n\nHello, World!", path.Base(tmpDir))
	assert.Equal(t, expected, res)
}

func TestMultiplePackagesCausesError(t *testing.T) {
	tmpDir := copyTestFileToGoPath(t, "bar/docs_main.go", "")
	defer os.RemoveAll(tmpDir)

	copyTestFileToGoPath(t, "foo/docs.go", tmpDir)

	_, err := ConvertDocs(tmpDir, DefaultConfig())
	assert.Error(t, err)
}

func TestExtraFilesCauseNoErrors(t *testing.T) {
	tmpDir := copyTestFileToGoPath(t, "foo/docs.go", "")
	defer os.RemoveAll(tmpDir)

	// Copy a random file in there too
	copyTestFileToGoPath(t, "extra.md", tmpDir)

	_, err := ConvertDocs(tmpDir, DefaultConfig())
	assert.NoError(t, err)
}

// copyTestFileToGoPath copies testFile from the testdata directory into a
// random temporary directory in the GOPATH. The returned string is the
// directory that was created. The caller should delete it when finished.
// To use an existing temp dir, pass a non-empty string as tmpDir.
func copyTestFileToGoPath(t *testing.T, testFile, tmpDir string) string {

	var err error
	if tmpDir == "" {
		tmpDir, err = ioutil.TempDir(path.Join(build.Default.GOPATH, "src"), "goreadme")
		require.NoError(t, err)
	}

	wd, err := os.Getwd()
	require.NoError(t, err)

	testFileBytes, err := ioutil.ReadFile(path.Join(wd, "testdata", testFile))
	require.NoError(t, err)

	err = ioutil.WriteFile(path.Join(tmpDir, path.Base(testFile)),
		testFileBytes, 0644)
	require.NoError(t, err)

	return tmpDir
}
