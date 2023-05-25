package main

import (
	"fmt"

	"github.com/wongearl/karmadactl-util/pkg/common"
	"github.com/wongearl/karmadactl-util/pkg/join"
	"github.com/wongearl/karmadactl-util/pkg/k8sclient"
	"github.com/wongearl/karmadactl-util/pkg/unjoin"
	"k8s.io/client-go/kubernetes"
)

func main() {
	client, err := k8sclient.NewKubernetesClient("/home/king/.kube/config", "", 30, 5)
	if err != nil {
		fmt.Println("NewKubernetesClient err:", err)
		return
	}
	testJoin(client)
	testUnJoin(client)
}

func testJoin(k8sClient kubernetes.Interface) {
	j := &join.Join{
		KarmdactlArg: common.KarmdactlArg{
			ClusterName:                  "member1",
			ClusterConfigSecretName:      "member-config",
			ClusterConfigSecretNamespace: "default",
			ClusterContext:               "kubernetes-admin@cluster.local",
			KarmadaConfigSecretName:      "karmada-kubeconfig",
			KarmadaConfigSecretNamespace: "karmada-system",
			KarmadaContext:               "karmada-apiserver",
		},
	}
	stdout, stderr, err := j.JoinCluster(k8sClient)
	if err != nil {
		fmt.Println("err:", err, "stderr:", stderr)
		return
	}
	fmt.Println("stdout:", stdout)
}

func testUnJoin(k8sClient kubernetes.Interface) {
	u := &unjoin.Unjoin{
		KarmdactlArg: common.KarmdactlArg{
			ClusterName:                  "member1",
			ClusterConfigSecretName:      "member-config",
			ClusterConfigSecretNamespace: "default",
			ClusterContext:               "kubernetes-admin@cluster.local",
			KarmadaConfigSecretName:      "karmada-kubeconfig",
			KarmadaConfigSecretNamespace: "karmada-system",
			KarmadaContext:               "karmada-apiserver",
		},
	}
	stdout, stderr, err := u.JoinCluster(k8sClient)
	if err != nil {
		fmt.Println("err:", err, "stderr:", stderr)
		return
	}
	fmt.Println("stdout:", stdout)
}
