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
	"hash/fnv"
	"os/user"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		cli, err := getDockerClient()
		if err != nil {
			return
		}

		containers, err := getContainers(cli)
		if err != nil {
			return
		}

		// format <user> <pid> <processname> <is_daemon=[1|0]>
		for _, con := range containers {
			printPsDoomCompatible(con.ID, firstOrEmpty(con.Names))
		}
	},
}

func init() {
	dockerCmd.AddCommand(psCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printPsDoomCompatible(containerId, containerName string) {
	currentUser, err := user.Current()

	if err != nil {
		return
	}
	mappedConName, err := mapStringToInt(containerId)
	if err != nil {
		return
	}
	fmt.Printf("%s %d %s 1\n", currentUser.Username, mappedConName, containerName)

}

func mapStringToInt(s string) (uint32, error) {

	var algorithm = fnv.New32a()
	_, err := algorithm.Write([]byte(s))
	if err != nil {
		return 0, err
	}

	return algorithm.Sum32(), nil
}

func firstOrEmpty(sa []string) string {
	if len(sa) < 1 {
		return ""
	}
	return sa[0]
}
