/*
Copyright © 2022 AT&T <rx7322@att.com>
* ============LICENSE_START=======================================================
* Author: Richard Masci (rx7322)
* ================================================================================
* Copyright (C) 2017 - 2018 AT&T Intellectual Property. All rights
* 						reserved.
* ================================================================================
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
* ============LICENSE_END=========================================================

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testmeCmd represents the testme command
var testmeCmd = &cobra.Command{
	Use:   "testme",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testme called")
	},
}

func init() {
	rootCmd.AddCommand(testmeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
