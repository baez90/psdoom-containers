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
	"context"
	"github.com/baez90/psdoom-containers/internal/pkg/hashing"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
	"strconv"
)

// killCmd represents the kill
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := getDockerClient()
		if err != nil {
			return
		}

		containers, err := getContainers(cli)

		for _, container := range containers {
			mappedName, err := hashing.MapStringToInt(container.ID)
			if err == nil && strconv.Itoa(int(mappedName)) == args[0] {
				ctx := context.Background()
				_ = cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{
					Force: true,
				})
			}
		}
	},
}

func init() {
	dockerCmd.AddCommand(killCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// killCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// killCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
