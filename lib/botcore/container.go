package botcore

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

func pullImage(image string, imageMap map[string]string) {
	zerolog.TimeFieldFormat = ""

	ctx := context.Background()
	client, err := client.NewEnvClient()

	log.Debug("Pulling image with this client:", client)

	if err != nil {
		log.Error("Recieved the following error:", err)
		panic(err)
	}

	out, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		log.Error("Recieved the following error:", err)
		panic(err)
	}

	log.Debug("Pulled the image:", image)

	defer out.Close()

	io.Copy(os.Stdout, out)
}

func runContainer(image string, cmd string) {
	ctx := context.Background()
	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := client.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   []string{cmd},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}

func stopContainer() {
	ctx := context.Background()
	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		if err := client.ContainerStop(ctx, container.ID, nil); err != nil {
			panic(err)
		}
		fmt.Println("Success")
	}
}
