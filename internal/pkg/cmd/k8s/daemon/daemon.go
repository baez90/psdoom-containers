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
	"github.com/baez90/psdoom-containers/internal/pkg/api/k8s"
	"github.com/baez90/psdoom-containers/internal/pkg/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var k8sDaemonCmd = &cobra.Command{
	Use:   "k8s-daemon",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Error("Do not use this command directly but one of its sub-commands")
		os.Exit(1)
	},
}

func init() {
	cmd.RootCmd.AddCommand(k8sDaemonCmd)

	if home := k8s.HomeDir(); home != "" {
		k8s.KubeConfigPathFlag = k8sDaemonCmd.Flags().String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		k8s.KubeConfigPathFlag = k8sDaemonCmd.Flags().String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}
