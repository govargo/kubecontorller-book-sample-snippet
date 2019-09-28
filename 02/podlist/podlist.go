package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os/user"
	"path/filepath"
)

func main() {
	u, _ := user.Current()
	defaultPath := filepath.Join(u.HomeDir, ".kube", "config")
	// set kubeconfig flag
	kubeconfig := flag.String("kubeconfig", defaultPath, "kubeconfig config file")
	flag.Parse()

	// retrieve kubeconfig
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// get clientset for kubernetes resources
	clientset, _ := kubernetes.NewForConfig(config)

	// Get list of pod objects
	pods, _ := clientset.CoreV1().Pods("").List(metav1.ListOptions{})

	// show pod object to stdout
	for i, pod := range pods.Items {
		fmt.Printf("[Pod Name %d]%s\n", i, pod.GetName())
	}
}
