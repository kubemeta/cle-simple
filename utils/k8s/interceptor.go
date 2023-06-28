package k8s

import (
	"context"

	"github.com/loggie-io/loggie/pkg/discovery/kubernetes/apis/loggie/v1beta1"
	"github.com/loggie-io/loggie/pkg/discovery/kubernetes/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InterceptorInterface interface {
	GetInterceptor(ctx context.Context, name string) (*v1beta1.Interceptor, error)
	CreateInterceptor(ctx context.Context, interceptor *v1beta1.Interceptor) error
}

type interceptor struct {
	kubeCli *versioned.Clientset
}

var _ InterceptorInterface = &interceptor{}

func (i *interceptor) CreateInterceptor(ctx context.Context, icp *v1beta1.Interceptor) error {
	_, err := i.kubeCli.LoggieV1beta1().Interceptors().Create(ctx, icp, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (i *interceptor) GetInterceptor(ctx context.Context, name string) (*v1beta1.Interceptor, error) {
	return i.kubeCli.LoggieV1beta1().Interceptors().Get(ctx, name, v1.GetOptions{})
}

func NewInterceptorInterfaceForConfig() (InterceptorInterface, error) {

	cli, err := GetClient()
	if err != nil {
		return nil, err
	}

	return &interceptor{
		kubeCli: cli,
	}, nil
}
