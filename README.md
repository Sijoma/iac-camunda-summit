# summit-iac-camunda


## Terraform example

1. Go to terraform folder

```bash
terraform init
terraform apply
kustomize build ../coffee-shop-kubernetes-deployment/ | kubectl apply -f -
```


## Crossplane example

Install `package/crds` first.

```bash
kustomize build crossplane/. | kubectl apply -f - 
```

```bash
kubectl wait --for=condition=Ready=true clusters coffee-shop-automation
```

```
kustomize build coffee-shop-deployment/ | kubectl apply -f -
```
