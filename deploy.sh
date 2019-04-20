#!/bin/bash

docker push "${IMAGE_NAME}:latest" && docker push "${IMAGE_NAME}:${version}"
echo "Hello World"
openssl aes-256-cbc -K $encrypted_e121f330d406_key -iv $encrypted_e121f330d406_iv
  -in cluster-config.yaml.enc -out kube/cluster-config.yaml -d
ls
pwd
kubectl --kubeconfig="kube/cluster-config.yaml" apply -f ./kube/deployment.yml
kubectl --kubeconfig="kube/cluster-config.yaml" apply -f ./kube/service.yml