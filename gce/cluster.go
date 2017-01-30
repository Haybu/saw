package gce

import (
	"net/http"
	"encoding/base64"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/api/container/v1"

	amdatu "bitbucket.org/amdatulabs/amdatu-kubernetes-go/client"

	"log"
)


func buildHttpClient(clstr *container.Cluster) (*http.Client, error) {
	clientCertificate, err := base64.StdEncoding.DecodeString(clstr.MasterAuth.ClientCertificate)
	if err != nil {
		return nil, err
	}
	clientKey, err := base64.StdEncoding.DecodeString(clstr.MasterAuth.ClientKey)
	if err != nil {
		return nil, err
	}
	certificate, err := tls.X509KeyPair(clientCertificate, clientKey)

	caCertificate, err := base64.StdEncoding.DecodeString(clstr.MasterAuth.ClusterCaCertificate)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCertificate)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      caCertPool,
	}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return client, nil
}


func NewKubernetesClient(name, project, zone string) (*amdatu.Client, error) {
	gce, err := NewGCEContainerClient()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	cluster, err := GetGCECluster(gce, project, zone, name)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	username := cluster.MasterAuth.Username
	password := cluster.MasterAuth.Password
	host := "https://" + cluster.Endpoint

	httpClient, err := buildHttpClient(cluster)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	client := amdatu.NewClientWithHttpClient(httpClient, host, username, password)

	//fmt.Printf("cluster IP is %s username is %s password is %s", host, username, password)
	return &client, nil
}


func NewKubernetesClientFromClstr(clstr *container.Cluster) (*amdatu.Client, error) {
	username := clstr.MasterAuth.Username
	password := clstr.MasterAuth.Password
	host := "https://" + clstr.Endpoint

	httpClient, err := buildHttpClient(clstr)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	client := amdatu.NewClientWithHttpClient(httpClient, host, username, password)

	//fmt.Printf("cluster IP is %s username is %s password is %s", host, username, password)
	return &client, nil
}
