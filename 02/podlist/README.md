# PodList サンプル

## 概要

Kubernetesクラスタの全てのNamespace上のPodの一覧を取得し、表示します

## 使い方

```
$ export GO111MODULE=on
$ go run podlist.go -kubeconfig <kubeconfig file>
# e.g.
$ go run podlist.go -kubeconfig $HOME/.kube/config
```
