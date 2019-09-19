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

var assignLicenseCommand = &cobra.Command{
	Use:    "assign-license [account_id] [license_id]",
	Short:  "Assign license for the given account ID and license ID",
	Run:    assignLicense,
	PreRun: prerunManager,
	Args:   cobra.ExactArgs(2),
}

func init() {
	RootCommand.AddCommand(assignLicenseCommand)
}

func assignLicense(cmd *cobra.Command, args []string) {
	uuid, err := manager.AssignLicense(args[0], args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error assigning license: %v\n", err)
		Exit(1)
	}
	fmt.Printf("Successfully assigned license '%v': %v\n", args[1], uuid)
}
