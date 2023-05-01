# summit-iac-camunda


## Terraform example

1. Go to terraform folder

```bash
terraform init
terraform apply
kustomize build ../coffee-shop-kubernetes-deployment/ | kubectl apply -f -
```


## Crossplane example (super experimental)

Install `package/crds` of camunda-crossplane-provider first and run it too! 

```bash
kustomize build crossplane/. | kubectl apply -f - 
```

```bash
kubectl wait --for=condition=Ready=true clusters coffee-shop-automation
```

```
kustomize build coffee-shop-deployment/ | kubectl apply -f -
```
