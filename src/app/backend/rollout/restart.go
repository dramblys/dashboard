package rollout

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClient "k8s.io/client-go/kubernetes"
	"time"
)

type RestartResult struct {
	Result *appsv1.Deployment `json:"result"`
}

func RestartDeployment(client k8sClient.Interface, namespace string, name string) (result *appsv1.Deployment, err error) {
	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	annotations := deployment.Spec.Template.Annotations
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	deployment.Spec.Template.Annotations = annotations

	deployment, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, v1.UpdateOptions{})

	return deployment, err
}
