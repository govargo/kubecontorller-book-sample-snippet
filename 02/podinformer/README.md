# PodInformer サンプル

## 概要

PodのEventを監視するInformerを作成し、Eventを標準出力に表示します

## 使い方

```
$ export GO111MODULE=on
$ go run podinformer.go -kubeconfig <kubeconfig file>
# e.g.
$ go run podinformer.go -kubeconfig $HOME/.kube/config
```

止める場合は「Ctrl + C」で止めてください
