package service

import (
	"context"
	"testing"
	"fmt"
	"time"

	"github.com/metrue/fx/api"
)

func createDockerfile(client api.FxServiceClient) (*api.UpMsgMeta, error) {
	ctx := context.Background()
	upReq := &api.UpDockerfileRequest{
		Dockerfiles: []*api.DockerfileMeta{
			&api.DockerfileMeta{
				Content: fmt.Sprintf(`FROM ubuntu \n
RUN apt-get install -y ruby rubygems git\n
RUN git clone https://github.com/luisbebop/docker-sinatra-hello-world.git /opt/sinatra/ \n
RUN gem install bundler \n
EXPOSE 3000 \n
RUN echo \"%d\"
RUN cd /opt/sinatra && git pull && bundle install \n
CMD [\"/usr/local/bin/foreman\",\"start\",\"-d\",\"/opt/sinatra\"]`, time.Now().Unix()),
			},
		},
	}

	upRes, err := client.UpDockerfile(ctx, upReq)
	if err != nil {
		return nil, err
	}

	if len(upRes.Instances) != 1 {
		return nil, fmt.Errorf("UpDockerfile response should have one instance, found %d", len(upRes.Instances))
	}
	if upRes.Instances[0].Error != "" {
		return nil, fmt.Errorf("Up error: %s", upRes.Instances[0].Error)
	}

	return upRes.Instances[0], nil
}

func TestUpDockerfile(t *testing.T) {
	runServer(t)

	client, conn, err := api.NewClient(grpcEndpoint)
	defer stopServer(conn)

	if err != nil {
		t.Fatal(err)
	}

	_, createErr := createDockerfile(client)
	if createErr != nil {
		t.Fatal("Up Dockerfile: %s\n", err.Error())
	}
}