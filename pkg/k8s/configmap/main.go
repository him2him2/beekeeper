package configmap

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Options represents configmap's options
type Options struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	Labels      map[string]string
	Data        map[string]string
	BinaryData  map[string][]byte
}

// Set creates ConfigMap, if ConfigMap already exists updates in place
func Set(ctx context.Context, clientset *kubernetes.Clientset, o Options) (err error) {
	spec := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:        o.Name,
			Namespace:   o.Namespace,
			Annotations: o.Annotations,
			Labels:      o.Labels,
		},
		BinaryData: o.BinaryData,
		Data:       o.Data,
	}

	_, err = clientset.CoreV1().ConfigMaps(o.Namespace).Create(ctx, spec, metav1.CreateOptions{})
	if err != nil {
		if !k8sErrors.IsNotFound(err) {
			fmt.Printf("configmap %s already exists in the namespace %s, updating the map\n", o.Name, o.Namespace)
			_, err = clientset.CoreV1().ConfigMaps(o.Namespace).Update(ctx, spec, metav1.UpdateOptions{})
			if err != nil {
				return err
			}
		}
		return err
	}

	return
}