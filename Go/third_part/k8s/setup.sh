#!/bin/bash

echo "先准备基础依赖，安装 helm"
brew install helm

echo "使用 helm 安装 k8s-dashboard"
# Add kubernetes-dashboard repository
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
# Deploy a Helm Release named "kubernetes-dashboard" using the kubernetes-dashboard chart
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard

echo "或者直接使用 kubectl 安装"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.6.1/aio/deploy/recommended.yaml
echo "访问 dashboard"
kubectl proxy
echo "现在，你可以打开网址 http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/ 进行访问了"


echo "接下来，我们要执行 k8s-admin.yaml 文件了，文件具体内容如下："
bat k8s-admin.yaml
sleep 2

echo "首先，可以分开执行"
echo "先执行 获取管理员角色的 secret 名称"
kubectl get secrets -n kube-system | grep dashboard-admin | awk '{print $1}'
echo "再执行 对应的管理员的 token 值"
kubectl describe secret dashboard-admin-token-dknqx -n kube-system

echo "其次，可以一起执行"
kubectl describe secret dashboard-admin-token-dknqx -n kube-system | grep -E '^token' | awk '{print $2}'
