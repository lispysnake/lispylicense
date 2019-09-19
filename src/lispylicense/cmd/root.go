//
// Copyright © 2019 Lispy Snake, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"license"
	"os"
)

var (
	manager *license.Manager
)

// RootCommand is the top-level command (and parent) of our CLI
var RootCommand = &cobra.Command{
	Use:   "lispylicense",
	Short: "lispylicense is the Lispy Snake, Ltd. license tool",
}

// Private helper for commands requiring a manager instance.
func prerunManager(cmd *cobra.Command, args []string) {
	m, err := license.NewManager()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	manager = m
}

// Close will help close up the CLI resources
func Close() {
	if manager == nil {
		return
	}
	manager.Close()
	manager = nil
}

// Exit is used to prevent leaking the Manager.
func Exit(statusCode int) {
	Close()
	os.Exit(statusCode)
}
