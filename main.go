package main

import (
	"fmt"

	"github.com/wongearl/karmadactl-util/pkg/join"
	"github.com/wongearl/karmadactl-util/pkg/k8sclient"
	"k8s.io/client-go/kubernetes"
)

func main() {
	client, err := k8sclient.NewKubernetesClient("/home/king/.kube/config", "", 30, 5)
	if err != nil {
		fmt.Println("NewKubernetesClient err:", err)
		return
	}
	testJoin(client)
}

func testJoin(k8sClient kubernetes.Interface) {
	j := &join.Join{
		ClusterName:                  "member1",
		ClusterConfigSecretName:      "member-config",
		ClusterConfigSecretNamespace: "default",
		ClusterContext:               "kubernetes-admin@cluster.local",
		KarmadaConfigSecretName:      "karmada-kubeconfig",
		KarmadaConfigSecretNamespace: "karmada-system",
		KarmadaContext:               "karmada-apiserver",
	}
	stdout, stderr, err := j.JoinCluster(k8sClient)
	if err != nil {
		fmt.Println("err:", err, "stderr:", stderr)
		return
	}
	fmt.Println("stdout:", stdout)
}
