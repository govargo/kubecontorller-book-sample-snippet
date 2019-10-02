package main

import (
	"flag"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
	"log"
	"path/filepath"
	"time"
)

func main() {
	var defaultKubeConfigPath string
	if home := homedir.HomeDir(); home != "" {
		// build kubeconfig path from $HOME dir
		defaultKubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	// set kubeconfig flag
	kubeconfig := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig config file")
	flag.Parse()
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create InformerFactory
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)

	// Create pod informer by informerFactory
	podInformer := informerFactory.Core().V1().Pods()

	// Create RateLimitQueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	// shutdown when process ends
	defer queue.ShutDown()

	// Add EventHandler to informer
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(old interface{}) {
			var key string
			var err error
			if key, err = cache.MetaNamespaceKeyFunc(old); err != nil {
				runtime.HandleError(err)
				return
			}
			queue.Add(key)
			log.Println("Added: " + key)
		},
		UpdateFunc: func(old, new interface{}) {
			var key string
			var err error
			if key, err = cache.MetaNamespaceKeyFunc(new); err != nil {
				runtime.HandleError(err)
				return
			}
			queue.Add(key)
			log.Println("Updated: " + key)
		},
		DeleteFunc: func(old interface{}) {
			var key string
			var err error
			if key, err = cache.MetaNamespaceKeyFunc(old); err != nil {
				runtime.HandleError(err)
				return
			}
			queue.Add(key)
			log.Println("Deleted: " + key)
		},
	})

	// Start Go routines
	informerFactory.Start(wait.NeverStop)
	// Wait until finish caching with List API
	informerFactory.WaitForCacheSync(wait.NeverStop)

	// Create Pod Lister
	podLister := podInformer.Lister()
	// Get List of pods
	_, err = podLister.List(labels.Nothing())
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second * 1)
	}
}
