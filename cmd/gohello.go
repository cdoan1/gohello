package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

var data_example = `
clusters:
  - name: example1
    url: https://icp-console.apps.scale1.dev.multicloudops.io
    username: example1
    password: example1
  - name: example2
    url: https://icp-console.apps.offline.dev.multicloudops.io
    url-ocp: https://console-openshift-console.apps.offline.dev.multicloudops.io
    username: exmaple2
    password: example2
  - name: blue-12
    url:
    username: kubeadmin
    password:
`

// Cluster is a struct of an array of string values
// * the fields have to start with capital letters
// * the `yaml:"name"` is what is searched in the yaml file to match
type Cluster struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Clusters is a group of cluster
type Clusters struct {
	Clusters []Cluster `yaml:"clusters"`
}

func main() {

	log.Println("hello world!")

	filename, _ := filepath.Abs("./data.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	check(err)
	log.Printf("Reading clusters from %s", string(filename))

	var cluster Clusters
	// err := yaml.Unmarshal([]byte(data_example), &cluster)
	err = yaml.Unmarshal([]byte(yamlFile), &cluster)
	check(err)

	// log.Printf("Clusters: %#v\n", cluster)
	d, err := yaml.Marshal(&cluster)
	log.Printf("current settings:\n\n%v\n", string(d))

	// log.Printf("length %d", len(cluster.Clusters))
	// log.Printf("length %s", cluster.Clusters[0].Name)
	// iterate through the clusters
	for _, element := range cluster.Clusters {
		if element.URL != "" && element.Password != "" {
			log.Printf("name \t%s", element.Name)
			log.Printf("url \t%s", element.URL)
			log.Printf("username \t%s", element.Username)
			log.Printf("password \t%s", element.Password)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
