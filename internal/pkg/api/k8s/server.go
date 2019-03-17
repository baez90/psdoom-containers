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

package k8s

import (
	"context"
	"fmt"
	"github.com/baez90/psdoom-containers/internal/pkg/api/k8s/generated"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

func NewK8sAPIServer(pr PodRegistry) k8sApi.K8SDaemonServer {
	client, err := GetKubeClient()

	if err != nil {
		panic(err)
	}
	return &k8sDaemonServer{
		podRegistry: pr,
		client:      client,
	}
}

type k8sDaemonServer struct {
	podRegistry PodRegistry
	client      *kubernetes.Clientset
}

func (ds *k8sDaemonServer) GetPods(context.Context, *k8sApi.Empty) (*k8sApi.Pods, error) {
	logrus.Info("Processing GetPods() request...")
	pods := make([]*k8sApi.Pod, 0)
	ds.podRegistry.ForEach(func(key string, pod v1.Pod) {
		pods = append(pods, &k8sApi.Pod{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Id:        key,
		})
	})

	return &k8sApi.Pods{
		Pods: pods,
	}, nil
}

func (ds *k8sDaemonServer) KillPod(ctx context.Context, podDel *k8sApi.PodDeletion) (*k8sApi.Empty, error) {
	logrus.Infof("Processing KillPod() request for ID %s", podDel.PodId)
	pod, exists := ds.podRegistry.GetPod(strconv.Itoa(int(podDel.PodId)))
	if !exists {
		return &k8sApi.Empty{}, fmt.Errorf("no pod with given id %s/%s found", pod.Namespace, pod.Name)
	}

	err := ds.client.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &v1meta.DeleteOptions{})
	if err != nil {
		return &k8sApi.Empty{}, err
	}

	return &k8sApi.Empty{}, nil
}
