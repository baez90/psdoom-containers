# psdoom-containers

![](assets/img/screen1.png)

## Why?

[psdoom-ng](https://github.com/orsonteodoro/psdoom-ng) is a "remake" of the classic Doom but as kind of a _"task manager"_.
Beside of the possibility to shoot at your own pc's processes it is also possible to add custom commands to load, renice and kill processes.

`psdoom-containers` is a tiny CLI that creates output in the expected structure of psdoom-ng:

```
<user> <pid> <processname> <is_daemon=[1|0]>
```

but based on Docker containers or Kubernetes pods (Podman support is planned).

## Installation

```bash
go get -u github.com/baez90/psdoom-containers/cmd/psdoom-containers
```

## Usage

### Docker

```bash
export PSDOOMPSCMD='psdoom-containers docker ps'
export PSDOOMRENICECMD='true'
export PSDOOMKILLCMD='psdoom-containers docker kill'
```

there's also the script `./scripts/psdoom-docker.sh` that is setting the environemnt variables and starting `psdoom-ng` with some default parameters.

To get some tiny containers running that you can shoot at run the script `./scripts/spawn-containers.sh`.
It is creating a few Alpine Linux containers that are sleeping for 3600s to keep the load as low as possible.

### Kubernetes

`psdoom-ng` expects the username of the process owner to be in one column of the output.
While this is not directly possible for Docker containers (beside of probably something like labels) and therefore every container is "running as the current user" to have some targets, the Kubernetes implementation differs here.

For Kubernetes the `username` column is mapped to the pod namespace.
Therefore if you don't want to add the flag to shoot at all pods, ensure that there's a Kubernetes namespace corresponding to your username:

```bash
kubectl create ns $(whoami)
```

and after that you can deploy some pods to the namespace to have some targets.

```bash
export PSDOOMPSCMD="psdoom-containers k8s ps"
export PSDOOMRENICECMD="true"
export PSDOOMKILLCMD="psdoom-containers k8s kill"
```

there's also the script `./scripts/psdoom-k8s.sh` that is setting the environment variables and starting `psdoom-ng` with some default paramters.

To get a few tiny pods running you can deploy the `sleep-deployment.yaml` in the `./scripts` directory.
The deployment is configured to start 10 replicas of the same Alpine Linux container as the corresponding Docker script only running a `sleep` for 3600s.