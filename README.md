```sh
export NAMESPACE=<your namespace>
kubectl create namespace $NAMESPACE
kubectl config set-context $(kubectl config current-context) --namespace=$NAMESPACE
```
