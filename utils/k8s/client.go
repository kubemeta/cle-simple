package k8s

import (
	"github.com/loggie-io/loggie/pkg/discovery/kubernetes/client/clientset/versioned"
	"k8s.io/client-go/rest"
)

func GetClient() (*versioned.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	cli, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
