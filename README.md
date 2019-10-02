# kubecontorller-book-sample-snippet
技術書典7で頒布した「実践入門 Kubernetesカスタムコントローラへの道」の本文中に掲載したサンプルコード用のリポジトリです

## Update History

2019/10/02 Modified way to set kubeconfig path flag [podlist.go](02/podlist/podlist.go#L9-L21)  

2019/10/02 Modified way to set kubeconfig path flag [podinformer.go](02/podinformer/podinformer.go#L11-L25)   

2019/10/02 Modified way to set kubeconfig path flag [enqueuePod.go](02/workqueue/enqueuePod.go#L12-L27)   

|Before|After|
|:---:|:---:|
|kubeconfig := flag.String("kubeconfig", "~/.kube/config", "kubeconfig config file")|kubeconfig := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig config file")|
