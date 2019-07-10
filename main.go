package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	kubesystemNamespace, err := clientset.CoreV1().Namespaces().Get("kube-system", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	if kubesystemNamespace.ObjectMeta.Labels == nil {
		kubesystemNamespace.ObjectMeta.Labels = make(map[string]string)
	}

	_, ok := kubesystemNamespace.ObjectMeta.Labels["name"]
	if !ok {
		fmt.Println("'name' label for 'kube-system' namespace is missing. Labeling namespace...")
		kubesystemNamespace.ObjectMeta.Labels["name"] = "kube-system"
		_, err := clientset.CoreV1().Namespaces().Update(kubesystemNamespace)
		if err != nil {
			panic(err.Error)
		}
		fmt.Println("'kube-system' namespace labeled with 'name' label")
	} else {
		fmt.Println("namespace 'kube-system' has proper labels.")
	}
}
