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

package main

import (
	"github.com/baez90/psdoom-containers/internal/pkg/cmd"
	_ "github.com/baez90/psdoom-containers/internal/pkg/cmd/docker"
	_ "github.com/baez90/psdoom-containers/internal/pkg/cmd/k8s"
	_ "github.com/baez90/psdoom-containers/internal/pkg/cmd/k8s/daemon"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

	cmd.Execute()
}
