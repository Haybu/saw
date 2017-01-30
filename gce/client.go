package gce

import (
	"log"
	"encoding/json"
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
	amdatu "bitbucket.org/amdatulabs/amdatu-kubernetes-go/client"
)

// -------- Cluster level

func CreateCluster(name string, project string , zone string, nodes int64) (string, error) {

	gce, err := NewGCEContainerClient()
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	err = CreateGCECluster(gce, project, zone, name, int64(nodes))
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	clstr, err := WaitForGCEClusterProvisioning(gce, project, zone, name)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	k8s, err := NewKubernetesClientFromClstr(clstr)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	url := k8s.Url;

	return url, err
}

func DeleteCluster(name string, project string, zone string) error {
	gce, err := NewGCEContainerClient()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	err = DeleteGCECluster(gce, project, zone, name)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

func ClusterInfo(name string, project string, zone string) (string, error) {
	gce, err := NewGCEContainerClient()
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	cluster, err := GetGCECluster(gce, project, zone, name)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	j, _ := json.Marshal(cluster)

	return string(j), nil
}

// ----------- Kubernetes level ---------------------

func getK8sClient(name string, project string, zone string) (*amdatu.Client, error) {
	k8s, err := NewKubernetesClient(name, project, zone)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return k8s, nil
}

func ClusterNodes(name string, project string, zone string) (*v1.NodeList, error)  {
	k8s, _ := getK8sClient(name, project, zone)
	nodes, err := k8s.ListNodes()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
        return nodes, nil
}

func CreateNameSpace(namespace string, clusterName string, project string, zone string) (*v1.Namespace, error) {
	k8s, _ := getK8sClient(clusterName, project, zone)
	return k8s.CreateNamespace(namespace)
}

func DeleteNameSpace(namespace string, clusterName string, project string, zone string) error {
	k8s, _ := getK8sClient(clusterName, project, zone)
	return k8s.DeleteNamespace(namespace)
}

func CreateRC(appName string, image string, version string, clusterName string, namespace string, project string, zone string, replicas int32) (string, error){

	labels := map[string]string{"version": version}

	rc := v1.ReplicationController{
		ObjectMeta: v1.ObjectMeta{Name: appName, Namespace: namespace, Labels: labels},
		Spec: v1.ReplicationControllerSpec{
			Selector: map[string]string{"name": appName},
			Replicas: &replicas,
			Template: &v1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: map[string]string{"name": appName},
				},

				Spec: v1.PodSpec{
					Containers: []v1.Container{{
						Name:  appName,
						Image: image,
					}},
				},
			},
		},
	}

	k8s, _ := getK8sClient(clusterName, project, zone)
	createdRc, err := k8s.CreateReplicationController(namespace, &rc)

	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return createdRc.Name, nil
}

func ListPods(namespace string, clusterName string, project string, zone string) (*v1.PodList, error) {
	k8s, _ := getK8sClient(clusterName, project, zone);
	pods, err := k8s.ListPods(namespace)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return pods, nil
}

func CreateService(namespace string, appname string, clusterName string, project string, zone string) error {
	//k8s, _ := getK8sClient(clusterName, project, zone);
	//k8s.CreateService()
	return nil
}
