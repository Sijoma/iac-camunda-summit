# The channel containing the most recent version of Zeebe.
data "camunda_channel" "alpha" {
  name = "Alpha"
}

# A cluster plan type for default trials.
data "camunda_cluster_plan_type" "trial" {
  name = "Trial Cluster"
}

# An available region
data "camunda_region" "europe" {
  name = "Belgium, Europe (europe-west1)"
}

resource "camunda_cluster" "coffee_shop" {
  name = "Coffee Shop Terraform created"

  channel    = data.camunda_channel.alpha.id
  generation = data.camunda_channel.alpha.default_generation_id
  region     = data.camunda_region.europe.id
  plan_type  = data.camunda_cluster_plan_type.trial.id
}

resource "camunda_cluster_client" "coffee_shop" {
  name       = "coffee_shop_tf"
  cluster_id = camunda_cluster.coffee_shop.id
}

resource "kubernetes_namespace" "coffee_shop_terraform" {
  metadata {
    name = "coffee-shop-terraform"
  }
}

resource "kubernetes_secret" "zeebe_client" {
  metadata {
    name      = "service-worker-client"
    namespace = kubernetes_namespace.coffee_shop_terraform.id
  }

  data = {
    ZEEBE_CLIENT_ID                = camunda_cluster_client.coffee_shop.zeebe_client_id
    ZEEBE_CLIENT_SECRET            = camunda_cluster_client.coffee_shop.secret
    ZEEBE_ADDRESS                  = camunda_cluster_client.coffee_shop.zeebe_address
    ZEEBE_AUTHORIZATION_SERVER_URL = camunda_cluster_client.coffee_shop.zeebe_authorization_server_url
  }
}
