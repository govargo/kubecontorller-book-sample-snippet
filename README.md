# kubecontorller-book-sample-snippet
インプレスR&D NextPublishing出版「実践入門 Kubernetesカスタムコントローラへの道」の本文中に掲載したサンプルコード用のリポジトリです
底本となった技術書典7で頒布したコードとも互換があります。

https://nextpublishing.jp/book/11389.html

## Update History

2019/12/10 Update go version, go.mod, go.sum for [podlist](02/podlist), [podinformer](02/podinformer), [workqueue](02/workqueue)

2019/12/09 Added CRD [apiextensions.k8s.io/v1 yaml](01) validation-crd/sample-crd-validation-ga.yaml, addprintcolumn-crd/sample-crd-printcolumn-ga.yaml, subresource-crd/sample-crd-status-ga.yaml, subresource-crd/sample-crd-scale-ga.yaml, defaulting-crd/ 

2019/10/02 Modified way to set kubeconfig path flag [podlist.go](02/podlist/podlist.go#L9-L21)  

2019/10/02 Modified way to set kubeconfig path flag [podinformer.go](02/podinformer/podinformer.go#L11-L25)   

2019/10/02 Modified way to set kubeconfig path flag [enqueuePod.go](02/workqueue/enqueuePod.go#L12-L27)   

|Before|After|
|:---:|:---:|
|kubeconfig := flag.String("kubeconfig", "~/.kube/config", "kubeconfig config file")|kubeconfig := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig config file")|

2019/10/03 Added comment to clientcmd.BuildConfigFromFlags [podlist.go](02/podlist/podlist.go#L24) 

|Before|After|
|:---:|:---:|
||// retrieve kubeconfig|
