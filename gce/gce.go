package gce

import (
	"google.golang.org/api/option"
	"time"
	"fmt"

	"google.golang.org/api/transport"
	"context"
	"google.golang.org/api/container/v1"

)

// newContainerClient creates a new client for the Google Container Engine API.
func NewGCEContainerClient() (*container.Service, error) {
	ctx := context.Background()
	o := []option.ClientOption{
		option.WithEndpoint("https://container.googleapis.com/"),
		// option.WithScopes(container.CloudPlatformScope),
	}
	httpClient, endpoint, err := transport.NewHTTPClient(ctx, o...)
	if err != nil {
		return nil, err
	}
	client, err := container.New(httpClient)
	if err != nil {
		return nil, err
	}
	client.BasePath = endpoint
	return client, nil
}

func CreateGCECluster(gce *container.Service, project string, zone string, cluster string, nodes int64) error {
	req := &container.CreateClusterRequest{}
	req.Cluster = &container.Cluster{Name: cluster, InitialNodeCount: nodes}
	_, err := gce.Projects.Zones.Clusters.Create(project, zone, req).Do()
	return err
}

func DeleteGCECluster(gce *container.Service, project string, zone string, cluster string) error {
	_, err := gce.Projects.Zones.Clusters.Delete(project, zone, cluster).Do()
	return err
}

func GetGCECluster(gce *container.Service, project string, zone string, cluster string) (*container.Cluster, error) {
	return gce.Projects.Zones.Clusters.Get(project, zone, cluster).Do()
}

func WaitForGCEClusterProvisioning(gce *container.Service, project string, zone string, cluster string) (*container.Cluster, error) {
	for {
		clstr, err := GetGCECluster(gce, project, zone, cluster)
		if err != nil {
			return nil, err
		}
		switch clstr.Status {
		case "PROVISIONING":
			time.Sleep(5 * time.Second)
		case "RUNNING":
			return clstr, nil
		default:
			return nil, fmt.Errorf("invalid cluster status: %s", clstr.Status)
		}
	}
}
