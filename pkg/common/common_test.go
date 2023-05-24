package common

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	testclient "k8s.io/client-go/kubernetes/fake"
)

var (
	// Scheme contains all types of custom clientset and kubernetes client-go clientset
	Scheme = runtime.NewScheme()
)

func TestWriteKubeconfig(t *testing.T) {

	byteSline := []byte(`apiVersion: v1
	kind: Config
	clusters:
	  - cluster:
		  certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURtVENDQW9HZ0F3SUJBZ0lVRmMvUTErOGEvV0hPSytBem90dEVpUnM4VVRnd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1hERUxNQWtHQTFVRUJoTUNlSGd4Q2pBSUJnTlZCQWdNQVhneENqQUlCZ05WQkFjTUFYZ3hDakFJQmdOVgpCQW9NQVhneENqQUlCZ05WQkFzTUFYZ3hDekFKQmdOVkJBTU1BbU5oTVJBd0RnWUpLb1pJaHZjTkFRa0JGZ0Y0Ck1CNFhEVEl6TURVeU5EQTJNekl5TmxvWERUSTBNRFV5TXpBMk16SXlObG93WERFTE1Ba0dBMVVFQmhNQ2VIZ3gKQ2pBSUJnTlZCQWdNQVhneENqQUlCZ05WQkFjTUFYZ3hDakFJQmdOVkJBb01BWGd4Q2pBSUJnTlZCQXNNQVhneApDekFKQmdOVkJBTU1BbU5oTVJBd0RnWUpLb1pJaHZjTkFRa0JGZ0Y0TUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGCkFBT0NBUThBTUlJQkNnS0NBUUVBeG5obnEzR3o3VHV3WFBqSTVLYmJ0U25tOCs2MXVrZGdjdjRWTE1taWV0NkwKb1pJQzZ0WTVNdkdDNkpSSjFGMXJGRVorNkJ0UVViWUxRUXYvUHVUV3FqWjJXVEFLVFZ5dFJCdUFyM0t0bUNLNgpidXBscnAvcTZnaHVNK1pnbjU4M0dWR2dNdHBtUlJBOWZzYmdNNkJqQWY1eVZ4TnZqVmRUeFZDdURrOWR2THVLCkMzMi9UTVNJUUNUTXVPb3RSMkIraEJHSkJiS2x6aFhDRXhHZktpT3hPMm52T3pEN2xhU0pGQTcwSTdiN3hQNXQKTFQ5T3JuVWMwOXU5YVVkMmxVVzJoRXRqbXZiVllzWSsxS01maDA3cmJXQ3hVN0FvN0xjazcxSVdpbzh6ZklDTQpFMEw0U3gxZmN3cG1NdlVzNXB2VkI1T3BwV29SVnUreExLUTNic0lnaFFJREFRQUJvMU13VVRBZEJnTlZIUTRFCkZnUVVvalVzc2FWSU96UEJrVDRJUmswaTMvRW1CRzR3SHdZRFZSMGpCQmd3Rm9BVW9qVXNzYVZJT3pQQmtUNEkKUmswaTMvRW1CRzR3RHdZRFZSMFRBUUgvQkFVd0F3RUIvekFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBQ0ozdgpiQkVmbXI3T1BDK0xKM2RhTjJBMTZEY0ZRZno4VlJRVTJ3OEFLRklFMUZrR3NYeitCbEdvNjlWTjVUV2MwK0xBCk0zd2VKVktjek1TVXlDUEhIb1JyK2tJSDFZT3F1Wlh3ZmtrRnpJVzBZZzFlU3FkYVhmWWQrMmNmNkJUVFFDbngKZXlodTVONU5hUDd6NkxBNGNxdVBxaGo2NmFxVzQwc2toN1ZkeTVkWEdkVkZCUnFTREppWXBsVC8rM0VCd1lDYwppeUFhTWl5OFJWaDhpZmx4MGhuamc2RXR3aStZSkwyTVJSWlhqS1A3VjJWM0JMVE56NmgyU09nRlFneFVaTTA0CmsrRGl3Z2Q1a3AzZnplUm5PRWl6UEZpNWk2a0VSMEY2Zll2N3pnWm9YUCtRS1dyNzh1SjFabHI3dS9UQW01UkQKV3NYcEh2QWJqSmxDQm5UQjF3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
		  insecure-skip-tls-verify: false
		  server: https://karmada-apiserver.karmada-system.svc.cluster.local:5443
		name: karmada-apiserver
	users:
	  - user:
		  client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVPVENDQXlHZ0F3SUJBZ0lVS2tzVWt4NHhlK0VzdzRORmluZ1I0dC9mT1RBd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1hERUxNQWtHQTFVRUJoTUNlSGd4Q2pBSUJnTlZCQWdNQVhneENqQUlCZ05WQkFjTUFYZ3hDakFJQmdOVgpCQW9NQVhneENqQUlCZ05WQkFzTUFYZ3hDekFKQmdOVkJBTU1BbU5oTVJBd0RnWUpLb1pJaHZjTkFRa0JGZ0Y0Ck1DQVhEVEl6TURVeU5EQTJNamN3TUZvWUR6SXhNak13TkRNd01EWXlOekF3V2pBd01SY3dGUVlEVlFRS0V3NXoKZVhOMFpXMDZiV0Z6ZEdWeWN6RVZNQk1HQTFVRUF4TU1jM2x6ZEdWdE9tRmtiV2x1TUlJQklqQU5CZ2txaGtpRwo5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBdC85KzFKNVhFU2t0OFYxNHphekVKdHhwUkFMZVd6ek1pNVpWCkt5VWhUYWpZVkluUHozKzVETWh4c1pZSkdVelgvM1FpL1k2VnROUTV1ODBsWEZ5TWxuT0lZV2RIUFNvZkpQYWMKM2RnbGp4RzRmZFdrbk0vcTJlS0VvVGlkZGoxd0tod1dkSm1zQXRJd0F2ZzhCc3VUWkVsV2JWNUQ5Mlh5V0lFYgp4dVFiWGlQbG5aRCtQN1BuNmlDOE1ncFhaRVlHU0JNcXF3dFU5dnQvK0xyc2Y3UlZudGg3MTBTN21FMmJMM0ViCjFvMWFDMUhIbVgxRThSU2g2aXZySkhvanl2QTI5ZnNXMUxUeS9JTFN6OGt0UmEvUEhnbDlZL0x2SmpTa2hydlgKNC92REZyMERpYXpoOUNxK1BqSEh1NmkvYmFrSG10cG4xQjRCR3VHTlZlOHZOUld2WFFJREFRQUJvNElCR3pDQwpBUmN3RGdZRFZSMFBBUUgvQkFRREFnV2dNQjBHQTFVZEpRUVdNQlFHQ0NzR0FRVUZCd01DQmdnckJnRUZCUWNECkFUQU1CZ05WSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJRNkZnWi81Vk42ODJaVjQwdnl4UkhZQTZoVWREQWYKQmdOVkhTTUVHREFXZ0JTaU5TeXhwVWc3TThHUlBnaEdUU0xmOFNZRWJqQ0Jsd1lEVlIwUkJJR1BNSUdNZ2hacgpkV0psY201bGRHVnpMbVJsWm1GMWJIUXVjM1pqZ2ljcUxtVjBZMlF1YTJGeWJXRmtZUzF6ZVhOMFpXMHVjM1pqCkxtTnNkWE4wWlhJdWJHOWpZV3lDSWlvdWEyRnliV0ZrWVMxemVYTjBaVzB1YzNaakxtTnNkWE4wWlhJdWJHOWoKWVd5Q0ZDb3VhMkZ5YldGa1lTMXplWE4wWlcwdWMzWmpnZ2xzYjJOaGJHaHZjM1NIQkg4QUFBRXdEUVlKS29aSQpodmNOQVFFTEJRQURnZ0VCQUlqY1FSamdqWS9kZVV5L0t4NkhrSmhDVmpJSzVjMGZRVC9FTjd0dHdUSkRBZHBQCjByWHRMVStYa1U0Wkt4cFgrM3psUlFOc3dNQUoyMEI1ZVJSS1N1QkZMVHVJVSs5S2V6MXRJd1diWlVuR0VSamEKdzZ1MDhlSDdiYWJQUVRlT0orMWZNUlk2UE1RWG1pTXV4Z2Q2RXYwU1FTUmluaHRFVEw1TWxFd3lVWEVHUzIxQgpUdG4xaVJXVW9idGs3MUVtb0N6ZXVQL1BYdWVWbkx4bkx1T3F4QXRjRkxpaGc1TE1TamlWZm0zeXVsdGpWQWJoCkZxWFdCclU0YklxcVFhbmR1NTR1eTdZNEJmU1NJU1J3Tk1KRzc2U0JudzROWVE1cjV3R3AzQ0Nya2puaGUyZ3IKZWthMjlFVVA3TzgzR3U0a3lES1dZK1J5UktrSFhGVWprV29mbTVNPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
		  client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBdC85KzFKNVhFU2t0OFYxNHphekVKdHhwUkFMZVd6ek1pNVpWS3lVaFRhallWSW5QCnozKzVETWh4c1pZSkdVelgvM1FpL1k2VnROUTV1ODBsWEZ5TWxuT0lZV2RIUFNvZkpQYWMzZGdsanhHNGZkV2sKbk0vcTJlS0VvVGlkZGoxd0tod1dkSm1zQXRJd0F2ZzhCc3VUWkVsV2JWNUQ5Mlh5V0lFYnh1UWJYaVBsblpEKwpQN1BuNmlDOE1ncFhaRVlHU0JNcXF3dFU5dnQvK0xyc2Y3UlZudGg3MTBTN21FMmJMM0ViMW8xYUMxSEhtWDFFCjhSU2g2aXZySkhvanl2QTI5ZnNXMUxUeS9JTFN6OGt0UmEvUEhnbDlZL0x2SmpTa2hydlg0L3ZERnIwRGlhemgKOUNxK1BqSEh1NmkvYmFrSG10cG4xQjRCR3VHTlZlOHZOUld2WFFJREFRQUJBb0lCQVFDaFVDVnc0UFV5ZldqagpFSERsMlE4TGh2ZmZBYWpTVXFaOXhxb2FybTNaT1N1WVNrNWYrL2xQNUxnUTJwcFZqUkpkeXdNV1M1aWl0ZUY3CjdlcFFaRzk1UkFjUVpreUZxbFV0d2V4YmJySFhFZEV1dVV5ZGtvZXl5SzVBN25MV2hCeS9QbXJOaFNEU1JGYUYKYy82a2NueGhVdzZyeWhaS1l4MnFURjcrNjJPM0RzTUNsSllQY0Njc0tlTmdLWmtXdkpabmpIclhHNU05ZExneQpwTUxUR1NZOGJHWDJqNEk0NStETFVMU2IwODlwMVMvdVJFTmhHUFZ1VW9ZbVlnNU9VN0o1VHpVRk9JK2FvTmNNClovcnNweHRYZmE3VWh6WE1Gb25ycDNEZzg5V0NIZXkvbUthNmMyK05rUWoyUFRuOWdZRUdPNkR5K1pHUGNCVk4KME96SlhObTlBb0dCQU5OWk5ad0ZhVTR4RGRQbVY0ak0xd0tzbHl6cSs4ZmxKZi90MytnVTR1bEltZjE2MU9VNQpPOG5TTHVLd2ZOQXBnVXBWbWFXM3NjY1BHZ1lIUVVVR3ZueEZHaEVlWFJJejc5S2c2Z2tNUzI3cUZrSHBDQ216CkRmNHNZVkt2ZjZ5Umd0c3ZsN1lVVDhKN2ZzUzd0NTR0Vlp3bnEza2J3RU9VVE9rODVwb1NoV0JUQW9HQkFON2YKQ1Q1Y1czVVh3eEk5VnFjVUlIc1FOYytRK1IyMFE1c3Q1TkU1WUJ2NWpzK3dCMm9uZVMyUXo2VlNVSmdtZEFpLwpTdHl3MHNuSkV3WktUbmFEeHNzbVNISk43aXdGMUx2MnVBSGZ0aWIwdDZ5V01HVmlJU05vd3RiVUFJUlV2ckVUClh6NGpHTi8yYnhrNklVWnFkSFdhN3JTZ2FUdXdFbWM4YkM3RjMzdVBBb0dCQUlDcWZNS2hZTHlqaklHR1o1LzIKNUtiU0g2N08xNzJZT1l3WGF3ckZQR2M5TmRKbFp4cXR2MEpjM1FKUTQ0dHUyVEZCNzZvOXJOTTgxR0Q3SmJjNgpKZGxOMEZLL28zV2pmTXRELzNiR3IxMjAwUndMSEZjV2xOdzZkSDE4TGtRR3loMWFXZ2dWVVlGYTRaQXZuOWVDClRlNGxFSFZJZWNJcWxMQWh6Vm5iRUt2VkFvR0JBSzhUOExQL1k0MHhkSGx3bkJDMlcrbXd5MFRhY3dnbG92SlYKZENuejg0OG1WVXpMMEpkUW1QMzFnMWt3dDhVK2QrcWpNMUQ0eXkzZStrWDN1M21ZZldMN0dQQktUZTVoU2tPNgpSY1NiRkFHNFBrMmkwalBpaVh2Q2dVUzMyQXdjY295eVZpQisrN3g4WDd3bWtSczY1MiszblF0aDlDa0NZUUtlClViSFFtWVg3QW9HQUltN29ndjBwMFlURk0waWpmamtxUXNUcnRlUnZqcU1EMlhPUnF5S3VOTEY5b2tWU212MzgKcy9xQ3JoTDBYZFN6QkgzaWoxcngwbjRDVEF4RjZjM0NnY3FQL3FRekRPZEV6YXN6WkJQQ2hMcldDYkRJTndVZwpoN05neHJ4SkNmK2NFWWNucXpZRXUrYWFWWkY1b0RRVFFlUU82Y0NLOE1EbmFueHF1R05xVitjPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
		name: karmada-apiserver
	contexts:
	  - context:
		  cluster: karmada-apiserver
		  user: karmada-apiserver
		name: karmada-apiserver
	current-context: karmada-apiserver`)

	type args struct {
		kubeconfigPath  string
		secretName      string
		secretNamespace string
		k8sClient       *testclient.Clientset
	}

	tests := []struct {
		name    string
		secret  *v1.Secret
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "write karmadaConfig",
			secret: &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "karmada-system",
					Name:      "karmada-kubeconfig",
				},
				Data: map[string][]byte{"kubeconfig": byteSline},
			},
			args: args{
				kubeconfigPath:  "/tmp/karmadaConfig",
				secretName:      "karmada-kubeconfig",
				secretNamespace: "karmada-system",
				k8sClient:       testclient.NewSimpleClientset(),
			},
			want:    "/tmp/karmadaConfig",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.k8sClient.CoreV1().Secrets(tt.args.secretNamespace).Create(context.TODO(), tt.secret, metav1.CreateOptions{})
			got, err := WriteKubeconfig(tt.args.kubeconfigPath, tt.args.secretName, tt.args.secretNamespace, tt.args.k8sClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteKubeconfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteKubeconfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
