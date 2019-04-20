
docker push "${IMAGE_NAME}:latest" && docker push "${IMAGE_NAME}:${version}"
./kubectl
echo "Hello World"
./kubectl --kubeconfig=kube/cluster-config.yaml apply -f ./kube/deployment.yml
./kubectl --kubeconfig=kube/cluster-config.yaml apply -f ./kube/service.yml