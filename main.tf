terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }

  backend "s3" {}
}

provider "digitalocean" {
  token = var.digitalocean_token
}

resource "digitalocean_kubernetes_cluster" "kube" {
  name    = var.cluster_name
  region  = "sgp1"
  version = var.cluster_version

  node_pool {
    name       = var.node_pool_name
    size       = var.node_pool_size
    node_count = var.node_pool_node_count
  }
}
