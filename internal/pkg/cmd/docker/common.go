package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	_ "github.com/docker/docker/reference"
)

func getDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts()
}

func getContainers(client *client.Client) ([]types.Container, error) {
	ctx := context.Background()
	return client.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
}
