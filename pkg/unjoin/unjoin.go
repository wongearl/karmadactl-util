package unjoin

import (
	"path/filepath"

	"github.com/wongearl/karmadactl-util/pkg/common"
	"k8s.io/client-go/kubernetes"
)

type Unjoin struct {
	common.KarmdactlArg
	//Delete cluster and secret resources even if resources in the cluster targeted for unjoin are not removedsuccessfully.
	Force bool
}

func (u *Unjoin) UnJoinCluster(k8sClient kubernetes.Interface) (stdout, stderr string, err error) {
	clusterKubeconfigPath, err := common.WriteKubeconfig("/tmp/clusterKubeconfig", u.ClusterConfigSecretName, u.ClusterConfigSecretNamespace, k8sClient)
	if err != nil {
		return stdout, stderr, err
	}
	karmadaKubeconfig, err := common.WriteKubeconfig("/tmp/karmadaKubeconfig", u.KarmadaConfigSecretName, u.KarmadaConfigSecretNamespace, k8sClient)
	if err != nil {
		return stdout, stderr, err
	}

	if u.Force {
		return common.ExecAtLocal(filepath.Join("karmadactl"), "unjoin", u.ClusterName, "--kubeconfig="+karmadaKubeconfig,
			"--karmada-context="+u.KarmadaContext, "--cluster-kubeconfig="+clusterKubeconfigPath, "--cluster-context="+u.ClusterContext, "--force=false")
	}

	return common.ExecAtLocal(filepath.Join("karmadactl"), "unjoin", u.ClusterName, "--kubeconfig="+karmadaKubeconfig,
		"--karmada-context="+u.KarmadaContext, "--cluster-kubeconfig="+clusterKubeconfigPath, "--cluster-context="+u.ClusterContext)
}
