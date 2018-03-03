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
	"strconv"

	"fmt"

	"github.com/spf13/cobra"
	"github.com/yetship/lauv2ex"
	"github.com/yetship/lauv2ex/http"
)

var query string
var queryById bool

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			member *v2ex.Member
			err    error
		)

		memberService := http.NewMemberService("https://www.v2ex.com/api")

		if queryById {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			member, err = memberService.GetMemberByID(id)
			if err != nil {
				panic(err)
			}
		} else {
			member, err = memberService.GetMemberByUsername(args[0])
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("id: %d\n", member.ID)
		fmt.Printf("name: %s\n", member.Username)
		fmt.Printf("page: %s\n", member.URL)
		fmt.Printf("status: %s\n", member.Status)
		fmt.Printf("avatar: %s\n", member.AvatarLarge)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	userCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "Name or id for query")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	userCmd.Flags().BoolVarP(&queryById, "id", "i", false, "Whether query by id")
}
