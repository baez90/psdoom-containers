// Copyright © 2019 Peter Kurfer peter.kurfer@googlemail.com
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

package docker

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var setupCmd = &cobra.Command{
	Use: "setup",
	Short: "setup psdoom-ng to use Docker containers as monsters",
	Long: `To use this setup command just run in a shell:
	eval $(psdoom-containers docker setup)
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("PSDOOMPSCMD='%s docker ps'\n", os.Args[0])
		fmt.Println("PSDOOMRENICECMD='true'")
		fmt.Printf("PSDOOMKILLCMD='%s docker kill'\n", os.Args[0])
	},
}

func init() {
	dockerCmd.AddCommand(setupCmd)
}