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

package cli

import (
	"fmt"
	"log"
	"openconsentia/hlsadmin/internal/repo/file"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hlsadmin",
	Short: "HLSAdmin is a Hyperledger Sawtooth orchestrator",
	Long: `HLSAdmin is an orchestrator to help you spin-up and manage a Hyperledger Sawtooth
based network.`,
}

func init() {

	cobra.OnInitialize(initApp)

	startCmd := createStartCmd()
	rootCmd.AddCommand(startCmd)
}

func initApp() {
	_, isContainer := os.LookupEnv("CONTAINER")
	settingLoc, err := file.SettingsLocation(isContainer)
	if err != nil {
		log.Fatalf("Unable to create settings location. Reason: %v", err)
	}
	settingsFile, err := file.InitSettingLoc(settingLoc, "settings.yaml")
	if err != nil {
		log.Fatalf("Unable to create settings file. Reason: %v", err)
	}
	log.Printf("%v created", settingsFile)
}

// Execute is the cli entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
