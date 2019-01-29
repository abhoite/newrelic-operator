/*
 * Copyright 2017-2018 IBM Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package get

import (
	"github.com/spf13/cobra"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many NewRelic resources.",
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GetCmd.PersistentFlags().String("foo", "", "A help for foo")
	GetCmd.PersistentFlags().StringP("output", "o", "json", "Output format. table/json/yaml are supported")

	GetCmd.PersistentFlags().StringP("type-condition", "t", "all", "Alert condition type. Only used for 'alertsconditions' command. all|conditions|synthetics|ext|plugin|nrql|infrastructure are supported")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
