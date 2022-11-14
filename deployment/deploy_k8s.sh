kubectl apply -f k8s/formatter.yaml
kubectl apply -f k8s/gomunkong-frontend.yaml

for LANG in "python" "golang" "java" "javascript"; do
    kubectl apply -f k8s/lang_servers-"$LANG".yaml
done
