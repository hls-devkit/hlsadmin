// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configutil

import (
	"io/ioutil"
	"os"
	"path"
)

var (
	configFileExists func(name string) bool
	configFileCreate func(name string, content []byte) error
)

func init() {
	configFileExists = func(name string) bool {
		_, err := os.Stat(name)
		if os.IsExist(err) {
			return true
		}
		return false
	}

	configFileCreate = func(filename string, content []byte) error {
		err := ioutil.WriteFile(filename, content, 0644)
		if err != nil {
			return err
		}
		return nil
	}
}

func newConfigurationFile(foldername string, filename string, content []byte) (string, error) {
	filenameFull := path.Join(foldername, filename)
	if !configFileExists(filenameFull) {
		configFileCreate(filenameFull, content)
	}
	return filenameFull, nil
}
