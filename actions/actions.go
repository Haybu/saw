package actions

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"github.com/Haybu/saw/gce"
	"log"
)

func ClusterCreateAction(c *cli.Context) error {

	fmt.Fprintf(c.App.Writer, "as if a cluster is created.\n")
	name := c.String("name")
	project := c.String("project")
	zone := c.String("zone")
	nodes := c.Int64("nodes")

	fmt.Fprintf(c.App.Writer, "cluster %v in project %v at zone %v will be created with %v number of nodes.\n", name, project, zone, nodes)

	return nil
}

func ClusterNodesAction(c *cli.Context) error {
	name := c.String("name")
	project := c.String("project")
	zone := c.String("zone")

         nodes, err := gce.ClusterNodes(name, project, zone)
	 if err != nil {
		log.Fatal(c.App.Writer, err.Error())
		return err
	 }

	fmt.Fprintf(c.App.Writer, "\ncluster %s has %d nodes as follows:\n\n", name, len(nodes.Items))
	for _, node := range nodes.Items {
		fmt.Fprintf(c.App.Writer, "%s\n", node.Name)
	}

	return nil
}
func ClusterInfoAction(c *cli.Context) error {
	name := c.String("name")
	project := c.String("project")
	zone := c.String("zone")

	info, err := gce.ClusterInfo(name, project, zone)
	if err != nil {
		log.Fatal(c.App.Writer, err.Error())
		return err
	}

	fmt.Fprintf(c.App.Writer, info)
	return nil
}

func ClusterDeleteAction(c *cli.Context) error {
	name := c.String("name")
	project := c.String("project")
	zone := c.String("zone")

	err := gce.DeleteCluster(name, project, zone)
	if err != nil {
		log.Fatal(c.App.Writer, err.Error())
		return err
	}

	return nil
}

func BuildApplicationAction(c *cli.Context) error {
	return nil;
}

func DeployApplicationAction(c *cli.Context) error {
	clusterName := c.String("name")
	project := c.String("project")
	zone := c.String("zone")
	appName := c.String("appname")
	version := c.String("version")
	replicas := c.Int64("replicas")
	namespace := c.String("namespace")
	image := c.String("image")

	rcname, err := gce.CreateRC(appName,image,version,clusterName,namespace,project,zone,int32(replicas))
	if err != nil {
		log.Fatal(c.App.Writer, err.Error())
		return err
	}

	fmt.Fprintf(c.App.Writer, "Replication controller %s is created.\n", rcname)

	return nil
}

func ListPodsAction(c *cli.Context) error {
	clusterName := c.String("name")
	project := c.String("project")
	zone := c.String("zone")
	namespace := c.String("namespace")

	pods, err := gce.ListPods(namespace, clusterName, project, zone)
	if err != nil {
		log.Fatal(c.App.Writer, err.Error())
		return err
	}

	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}

	return nil
}