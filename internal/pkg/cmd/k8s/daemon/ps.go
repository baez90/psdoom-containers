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

package daemon

import (
	"context"
	"fmt"
	"github.com/baez90/psdoom-containers/internal/pkg/api/k8s/generated"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		con, err := grpc.Dial("127.0.0.1:1357", grpc.WithInsecure())
		if err != nil {
			return
		}
		client := k8sApi.NewK8SDaemonClient(con)
		pods, err := client.GetPods(context.Background(), &k8sApi.Empty{})
		if err != nil {
			return
		}

		for _, pod := range pods.Pods {
			fmt.Printf("%s %s %s 1\n", pod.Namespace, pod.Id, pod.Name)
		}
	},
}

func init() {
	k8sDaemonCmd.AddCommand(psCmd)
}
