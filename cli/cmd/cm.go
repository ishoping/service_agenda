// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"service_agenda/cli/req"
	"service_agenda/entity"

	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("participators")
		startDate, _ := cmd.Flags().GetString("startDate")
		endDate, _ := cmd.Flags().GetString("endDate")

		mt := entity.GetMeeting(name, participators, startDate, endDate, title)

		req.MeetingPost(mt)

	},
}

func init() {
	RootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("name", "n", "Anonymous", "提供创建者名字")
	cmCmd.Flags().StringP("title", "t", "Anonymous", "提供会议主题名字")
	cmCmd.Flags().StringArrayP("participators", "p", nil, "提供一个会议的参与者")
	cmCmd.Flags().StringP("startDate", "s", "Anonymous", "会议开始时间，格式为YYYY-MM-DD HH—MM,")
	cmCmd.Flags().StringP("endDate", "e", "Anonymous", "会议结束时间，格式为YYYY-MM-DD HH—MM,")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
