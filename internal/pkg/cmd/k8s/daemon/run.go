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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	v1meta "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"net"
	"os"
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

		logger := logrus.New()
		grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(logger.WriterLevel(logrus.InfoLevel), logger.WriterLevel(logrus.WarnLevel), logger.WriterLevel(logrus.ErrorLevel), 0))

		server := grpc.NewServer()
		k8sApi.RegisterK8SDaemonServer(server, k8s.NewK8sAPIServer(podRegistry))

		err = server.Serve(lis)
		if err != nil {
			logger.Error("Failed to start k8s-daemon", err)
		} else {
			logger.Info("Successfully started server...")
		}
	},
}

func init() {
	k8sDaemonCmd.AddCommand(runDaemonCmd)
}

func initRegistry(pr k8s.PodRegistry) {
	client, err := k8s.GetKubeClient()

	if err != nil {
		logrus.Error("failed to get K8s client", err)
		os.Exit(1)
	}

	pods, err := client.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		logrus.Error("Failed to fetch pods", err)
		os.Exit(1)
	}

	for _, pod := range pods.Items {
		pr.AddPod(pod)
	}
}

func watchPodEvents(pr k8s.PodRegistry) {

	client, err := k8s.GetKubeClient()

	if err != nil {
		logrus.Error("Failed to get K8s client", err)
		os.Exit(1)
	}

	watcher, err := client.CoreV1().Pods("").Watch(v1.ListOptions{
		Watch:                true,
		IncludeUninitialized: false,
	})
	if err != nil {
		logrus.Error("failed to create watch for pods", err)
		os.Exit(1)
	}

	for in := range watcher.ResultChan() {
		pod, ok := in.Object.(*v1meta.Pod)
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
