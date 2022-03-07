/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "A command which outputs the sum total of inputs",
	Long: `
	A command that takes in user inputs in integer values and outputs a total sum, 
	including multiple values which is not pre-determined (adaptive). For example:

... sum 1 2 3 4 
output is 10.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kannan cobra-cli calculator v1.0 running")
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
