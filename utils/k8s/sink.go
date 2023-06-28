package k8s

import (
	"context"

	"github.com/loggie-io/loggie/pkg/discovery/kubernetes/apis/loggie/v1beta1"
	"github.com/loggie-io/loggie/pkg/discovery/kubernetes/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SinkInterface interface {
	GetSink(ctx context.Context, name string) (*v1beta1.Sink, error)
	CreateSink(ctx context.Context, sink *v1beta1.Sink) error
}

type sink struct {
	kubeCli *versioned.Clientset
}

var _ SinkInterface = &sink{}

func (s *sink) CreateSink(ctx context.Context, sink *v1beta1.Sink) error {
	_, err := s.kubeCli.LoggieV1beta1().Sinks().Create(ctx, sink, v1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *sink) GetSink(ctx context.Context, name string) (*v1beta1.Sink, error) {
	sink, err := s.kubeCli.LoggieV1beta1().Sinks().Get(ctx, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return sink, nil
}

func NewSinkInterfaceForConfig() (SinkInterface, error) {
	cli, err := GetClient()
	if err != nil {
		return nil, err
	}

	return &sink{
		kubeCli: cli,
	}, nil
}
