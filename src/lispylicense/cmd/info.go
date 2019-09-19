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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var infoCommand = &cobra.Command{
	Use:    "info [name]",
	Short:  "Show information/stats for the given license",
	Run:    info,
	PreRun: prerunManager,
	Args:   cobra.ExactArgs(1),
}

var emitJSON bool

func init() {
	emitJSON = false
	infoCommand.PersistentFlags().BoolVarP(&emitJSON, "json", "j", false, "Emit JSON output")
	RootCommand.AddCommand(infoCommand)
}

func info(cmd *cobra.Command, args []string) {
	info, err := manager.GetInfo(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "No such license '%v': %v\n", args[0], err)
		Exit(1)
	}

	if emitJSON {
		j, err := json.MarshalIndent(&info, "", "\t")
		if err != nil {
			fmt.Fprintf(os.Stderr, "JSON Error: %v\n", err)
			Exit(1)
		}
		fmt.Println(string(j))
	} else {
		fmt.Printf("Information for license: %v\n", args[0])
		if info.MaxUsers > 0 {
			fmt.Printf("Maximum users: %v\n", info.MaxUsers)
			fmt.Printf("Remaining users: %v\n", info.RemainingUsers)
			fmt.Printf("Current users: %v\n", info.CurrentUsers)
			fmt.Printf("Fill ratio: %.2f%%\n", info.FillRatio*100.0)
		} else {
			fmt.Printf("Current users: %v\n", info.CurrentUsers)
		}
		fmt.Printf("License description: %s\n", info.Description)
	}
}
