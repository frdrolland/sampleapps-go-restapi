/*
Copyright © 2025 Frédéric ROLLAND <frolland.work@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Run:           executeVersionCommand,
}

var (
	jsonMode bool = false // enable/disable debug mode
)

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Flag for JSON output activation ; persistent, so that subcommands (if any) will inherit this flag
	versionCmd.PersistentFlags().BoolVarP(&jsonMode, "json", "j", false, "Print output in JSON format")
}

func executeVersionCommand(cmd *cobra.Command, args []string) {
	if jsonMode {
		slog.Debug("json output mode enabled")
	}

	version := "0.0.1-alpha1"

	if jsonMode == true {
		fmt.Printf(`{"product-version":"%s"}`, version)
	} else {
		fmt.Printf(`product-version: "%s"`, version)
	}

}
