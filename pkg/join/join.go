package join

import (
	"path/filepath"

	"github.com/wongearl/karmadactl-util/pkg/common"
	"k8s.io/client-go/kubernetes"
)

type Join struct {
	ClusterName                  string
	ClusterConfigSecretName      string
	ClusterConfigSecretNamespace string
	KarmadaConfigSecretName      string
	KarmadaConfigSecretNamespace string
	KarmadaContext               string
}

func JoinCluster(join Join, k8sClient kubernetes.Interface) (stdout, stderr []byte, err error) {
	clusterKubeconfigPath, err := common.WriteKubeconfig("/tmp/clusterKubeconfig", join.ClusterConfigSecretName, join.ClusterConfigSecretNamespace, k8sClient)
	if err != nil {
		return nil, nil, err
	}
	karmadaKubeconfig, err := common.WriteKubeconfig("/tmp/karmadaKubeconfig", join.KarmadaConfigSecretName, join.KarmadaConfigSecretNamespace, k8sClient)
	if err != nil {
		return nil, nil, err
	}
	return common.ExecAtLocal(filepath.Join("karmadactl"), "join", join.ClusterName, "--kubeconfig="+karmadaKubeconfig,
		"--karmada-context="+join.KarmadaContext, "--cluster-kubeconfig="+clusterKubeconfigPath)
}
