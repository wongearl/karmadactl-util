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

func (j *Join) JoinCluster(k8sClient kubernetes.Interface) (stdout, stderr string, err error) {
	clusterKubeconfigPath, err := common.WriteKubeconfig("/tmp/clusterKubeconfig", j.ClusterConfigSecretName, j.ClusterConfigSecretNamespace, k8sClient)
	if err != nil {
		return stdout, stderr, err
	}
	karmadaKubeconfig, err := common.WriteKubeconfig("/tmp/karmadaKubeconfig", j.KarmadaConfigSecretName, j.KarmadaConfigSecretNamespace, k8sClient)
	if err != nil {
		return stdout, stderr, err
	}
	return common.ExecAtLocal(filepath.Join("karmadactl"), "join", j.ClusterName, "--kubeconfig="+karmadaKubeconfig,
		"--karmada-context="+j.KarmadaContext, "--cluster-kubeconfig="+clusterKubeconfigPath)
}
