package k8sclient

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubernetesClient(kubeconfig, token string, qps float32, burst int) (client kubernetes.Interface, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	config.QPS = qps
	config.Burst = burst
	if token != "" {
		// override the token
		config = &rest.Config{
			Host: config.Host,
			TLSClientConfig: rest.TLSClientConfig{
				Insecure: true,
			},
			BearerToken: token,
		}
	}
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return
}
