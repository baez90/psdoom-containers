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
	"github.com/baez90/psdoom-containers/internal/pkg/hashing"
	"k8s.io/api/core/v1"
	"strconv"
	"sync"
)

type ForEachHandler func(key string, pod v1.Pod)

type PodRegistry interface {
	AddPod(pod v1.Pod)
	GetPod(key string) (v1.Pod, bool)
	RemovePod(pod v1.Pod)
	ForEach(handler ForEachHandler)
}

type podRegistry struct {
	mux  sync.Mutex
	pods map[string]v1.Pod
}

func NewPodRegistry() PodRegistry {
	return &podRegistry{
		mux:  sync.Mutex{},
		pods: make(map[string]v1.Pod),
	}
}

func (pr *podRegistry) AddPod(pod v1.Pod) {
	pr.mux.Lock()
	defer pr.mux.Unlock()

	mappedName, err := hashing.MapStringToInt(string(pod.UID))
	if err != nil {
		return
	}

	pr.pods[strconv.Itoa(int(mappedName))] = pod
}

func (pr *podRegistry) GetPod(key string) (v1.Pod, bool) {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	pod, exists := pr.pods[key]
	return pod, exists
}

func (pr *podRegistry) RemovePod(pod v1.Pod) {
	pr.mux.Lock()
	defer pr.mux.Unlock()

	mappedName, err := hashing.MapStringToInt(string(pod.UID))
	if err != nil {
		return
	}

	delete(pr.pods, strconv.Itoa(int(mappedName)))
}

func (pr *podRegistry) ForEach(handler ForEachHandler) {
	pr.mux.Lock()
	defer pr.mux.Unlock()

	for key, pod := range pr.pods {
		handler(key, pod)
	}
}
