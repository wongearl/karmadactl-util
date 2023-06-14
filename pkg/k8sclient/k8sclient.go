package k8sclient

import (
	"fmt"

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
	//config.Host =  config.Host + fmt.Sprintf("/apis/cluster.karmada.io/v1alpha1/clusters/%s/proxy", clusterName)
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return
}

func NewMemberKubernetesClient(kubeconfig, token, clusterName string, qps float32, burst int) (client kubernetes.Interface, err error) {
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
	config.Host = config.Host + fmt.Sprintf("/apis/cluster.karmada.io/v1alpha1/clusters/%s/proxy", clusterName)
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return
}
