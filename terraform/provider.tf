variable "camunda_client_id" {}
variable "camunda_client_secret" {}

terraform {
  required_providers {
    camunda = {
      source = "multani/camunda"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0.0"
    }
  }
}

provider "camunda" {
  client_id     = var.camunda_client_id
  client_secret = var.camunda_client_secret
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}
