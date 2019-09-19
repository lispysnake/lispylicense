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
	"os"
)

var (
	description string
	maxUsers    int
)

var createLicenseCommand = &cobra.Command{
	Use:    "create-license [name]",
	Short:  "Create a new license within the database",
	Run:    createLicense,
	PreRun: prerunManager,
	Args:   cobra.ExactArgs(1),
}

func init() {
	createLicenseCommand.PersistentFlags().IntVarP(&maxUsers, "max-users", "u", -1, "Set the maximum users for the license")
	createLicenseCommand.PersistentFlags().StringVarP(&description, "description", "d", "License", "Set the license description")
	RootCommand.AddCommand(createLicenseCommand)
}

func createLicense(cmd *cobra.Command, args []string) {
	if err := manager.CreateLicense(args[0], maxUsers, description); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating license: %v\n", err)
		Exit(1)
	}
	fmt.Printf("Successfully created license '%v'\n", args[0])
}
