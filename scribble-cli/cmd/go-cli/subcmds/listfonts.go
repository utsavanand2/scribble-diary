/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

// listfontsCmd represents the listfonts command
var listfontsCmd = &cobra.Command{
	Use:   "listfonts",
	Short: "Lists the fonts available on the server.",
	Long:  `Lists all of the fonts available on the server to use for caption on the image.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listfonts called")
	},
}

func init() {
	rootCmd.AddCommand(listfontsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listfontsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listfontsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
