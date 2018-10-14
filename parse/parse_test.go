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
	"io/ioutil"
	"os"
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
	wd, err := os.Getwd()
	require.NoError(t, err)

	expectedBytes, err := ioutil.ReadFile(expectedFile)
	require.NoError(t, err)

	out, err := ConvertDocs(wd+"/foo", config)

	assert.NoError(t, err)
	assert.Equal(t, string(expectedBytes), out)
}
