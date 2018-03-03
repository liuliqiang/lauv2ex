// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yetship/lauv2ex/http"
)

// siteCmd represents the site command
var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		subCmd := "stat"
		if len(args) > 0 {
			if args[0] != "stat" && args[0] != "info" {
				fmt.Println("Invalid params")
				return
			} else {
				subCmd = args[0]
			}
		}

		siteService := http.NewRawV2exService("https://www.v2ex.com/api")

		if subCmd == "stat" {
			stats, err := siteService.GetStats()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Max Topic: %d\n", *stats.TopicMax)
			fmt.Printf("Max User: %d\n", *stats.MemberMax)
		} else {
			info, err := siteService.GetInfo()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Doamin: %s\n", info.Domain)
			fmt.Printf("Title: %s\n", info.Title)
			fmt.Printf("Slogan: %s\n", info.Slogan)
			fmt.Printf("Description: %s\n", info.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(siteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// siteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// siteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
