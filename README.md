# iac-camunda-summit

This repo contains the following:

- coffee-shop-app: Camunda Service worker that deploys a mounted BPMN process
- coffee-shop-deployment: k8s deployment manifests & BPMN file & Crossplane Zeebe Client
- crossplane: manifests to create a c8 saas cluster
- terraform: creates c8 saas cluster, client, k8s secret for client secret

## Terraform example

1. Go to terraform folder

```bash
terraform init
terraform apply
kustomize build ../coffee-shop-kubernetes-deployment/ | kubectl apply -f -
```


## Crossplane example (super experimental)

Install `package/crds` of camunda-crossplane-provider first and run it too! 

- Comment
```bash
kustomize build crossplane/. | kubectl apply -f - 
```

```bash
kubectl wait --for=condition=Ready=true clusters coffee-shop-automation
```

```
kustomize build coffee-shop-deployment/ | kubectl apply -f -
```
