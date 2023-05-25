package common

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ExecAtLocal(cmd string, args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	command := exec.Command(cmd, args...)
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()
	return stdout.String(), stderr.String(), err
}

func WriteKubeconfig(kubeconfigPath, secretName, secretNamespace string, k8sClient kubernetes.Interface) (string, error) {
	secret, err := k8sClient.CoreV1().Secrets(secretNamespace).Get(context.TODO(), secretName, v1.GetOptions{})
	if err != nil {
		return kubeconfigPath, err
	}

	kubecofigStr := secret.Data["kubeconfig"]
	if Exists(kubeconfigPath) {
		if err = os.Remove(kubeconfigPath); err != nil {
			return kubeconfigPath, err
		}

	}
	file, err := os.Create(kubeconfigPath)
	if err != nil {
		return kubeconfigPath, err
	}
	if _, err = io.WriteString(file, string(kubecofigStr)); err != nil {
		return kubeconfigPath, err
	}
	return kubeconfigPath, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
