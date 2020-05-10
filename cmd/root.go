/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
)

var (
	ssmName   []string
	abort     bool
	overwrite bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ssm-loader",
	Short: "Load AWS ssm params and set them as system environment variables.",
	Long:  `Load AWS ssm params and set them as system environment variables.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(ssmName) != 0 {
			return printVersion()
		}
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().StringSliceVarP(&ssmName, "ssm-name", "n", nil, "AWS ssm names seperated by a ','")
	rootCmd.Flags().BoolVarP(&abort, "abort-if-duplicates", "a", false, "Exit if there is an existing env var.")
	rootCmd.Flags().BoolVarP(&overwrite, "overwrite-if-duplicates", "o", false, "Overwtie existing env var with identical name")
}
