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
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print program version and exit",
	Run:   version,
}

func init() {
	RootCommand.AddCommand(versionCommand)
}

func version(cmd *cobra.Command, args []string) {
	fmt.Printf("lispylicense %v\n\nCopyright © 2019 Lispy Snake, Ltd.\n", license.Version)
	fmt.Printf("Licensed under the Apache License, Version 2.0\n")
}
