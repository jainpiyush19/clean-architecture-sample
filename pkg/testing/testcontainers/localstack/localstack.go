package localstack

import (
	"context"
	"strconv"
	"strings"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Service type
type Service string

const (
	// DynamoDB Service
	DynamoDB Service = "dynamodb"
)

// ContainerBuilder ...
type ContainerBuilder struct {
	edgePort int
	context  context.Context
	services []Service
	networks []string
}

// Container is a representation of the actual running container
type Container struct {
	testcontainers.Container
	context  context.Context
	edgePort int
}

// Endpoint returns the endpoint of the container
func (c Container) Endpoint() (string, error) {
	host, err := c.Container.Host(c.context)

	if err != nil {
		return "", err
	}

	port, err := c.MappedPort(c.context, nat.Port(strconv.Itoa(c.edgePort)))
	if err != nil {
		return "", err
	}

	return "http://" + host + ":" + port.Port(), nil
}

// NewContainerBuilder creates a builder for your container
func NewContainerBuilder() ContainerBuilder {
	return ContainerBuilder{
		edgePort: 4566,
		context:  context.Background(),
	}
}

// WithServices appends a list of sercices you want to enable
func (b ContainerBuilder) WithServices(services ...Service) ContainerBuilder {
	b.services = append(b.services, services...)
	return b
}

// WithNetworks appends a list of networks you want the container to join
func (b ContainerBuilder) WithNetworks(networks ...string) ContainerBuilder {
	b.networks = append(b.networks, networks...)
	return b
}

// WithEdgePort defines the edge port
func (b ContainerBuilder) WithEdgePort(port int) ContainerBuilder {
	b.edgePort = port
	return b
}

// Build creates and starts the container
func (b ContainerBuilder) Build() (Container, error) {
	edgePort := strconv.Itoa(b.edgePort)

	var services []string

	for _, v := range b.services {
		services = append(services, string(v))
	}

	req := testcontainers.ContainerRequest{
		Image:        "localstack/localstack",
		ExposedPorts: []string{edgePort},
		Env: map[string]string{
			"SERVICES":  strings.Join(services, ","),
			"DEBUG":     "0",
			"EDGE_PORT": edgePort,
		},
		WaitingFor: wait.ForListeningPort(nat.Port(edgePort)),
	}
	c, err := testcontainers.GenericContainer(b.context, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return Container{}, err
	}

	return Container{
		Container: c,
		context:   b.context,
		edgePort:  b.edgePort,
	}, nil
}

// WithContext allows the builder to utilize specified context
func (b ContainerBuilder) WithContext(ctx context.Context) ContainerBuilder {
	b.context = ctx
	return b
}
