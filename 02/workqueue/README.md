# WorkQueue Enqueue サンプル

## 概要

WorkQueueを生成し、WorkQueueにPod ObjectをEnqueueします

## 使い方

```
$ export GO111MODULE=on
$ go run enqueuePod.go -kubeconfig <kubeconfig file>
# e.g.
$ go run enqueuePod.go -kubeconfig $HOME/.kube/config
```

止める場合は「Ctrl + C」で止めてください
