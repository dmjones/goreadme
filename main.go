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
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dmjones/goreadme/parse"

	"github.com/BurntSushi/toml"
)

const configFile = ".goreadme.toml"

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Usage: goreadme [out-file]")
	}

	config, err := readConfig()
	logFatal(err, "Failed to read config")

	wd, err := os.Getwd()
	logFatal(err, "Failed to get directory info")

	output, err := parse.ConvertDocs(wd, config)
	logFatal(err, "Failed to parse package docs")

	if len(os.Args) == 1 {
		fmt.Println(output)
	} else {
		err = ioutil.WriteFile(os.Args[1], []byte(output), 0644)
		logFatal(err, "Failed to write to file")
	}
}

// readConfig reads the config file, if present, or returns the default config.
func readConfig() (*parse.Config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return parse.DefaultConfig(), nil
	}

	// Begin with default config. User can then overwrite what they want.
	c := parse.DefaultConfig()
	_, err := toml.DecodeFile(configFile, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func logFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg + ": " + err.Error())
	}
}
