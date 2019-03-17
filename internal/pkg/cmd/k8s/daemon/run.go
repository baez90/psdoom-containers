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
	"fmt"
	"github.com/baez90/psdoom-containers/internal/pkg/api/k8s"
	"github.com/baez90/psdoom-containers/internal/pkg/api/k8s/generated"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"net"
)

var runDaemonCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		podRegistry := k8s.NewPodRegistry()
		go watchPodEvents(podRegistry)
		go initRegistry(podRegistry)

		lis, err := net.Listen("tcp", "127.0.0.1:1357")

		if err != nil {
			fmt.Println(err)
			return
		}

		server := grpc.NewServer()
		k8sApi.RegisterK8SDaemonServer(server, k8s.NewK8sAPIServer(podRegistry))

		err = server.Serve(lis)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	k8sDaemonCmd.AddCommand(runDaemonCmd)
}

func initRegistry(pr k8s.PodRegistry) {
	client, err := k8s.GetKubeClient()

	if err != nil {
		panic(err)
	}

	pods, err := client.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		return
	}

	for _, pod := range pods.Items {
		pr.AddPod(pod)
	}
}

func watchPodEvents(pr k8s.PodRegistry) {

	client, err := k8s.GetKubeClient()

	if err != nil {
		panic(err)
	}

	watcher, err := client.CoreV1().Pods("").Watch(v1.ListOptions{
		Watch:                true,
		IncludeUninitialized: false,
	})
	if err != nil {
		fmt.Println(err)
	}

	for in := range watcher.ResultChan() {
		pod, ok := in.Object.(*v12.Pod)
		if !ok {
			continue
		}

		switch in.Type {
		case watch.Added:
			pr.AddPod(*pod)
			break
		case watch.Deleted:
			pr.RemovePod(*pod)
			break
		}
	}
}
