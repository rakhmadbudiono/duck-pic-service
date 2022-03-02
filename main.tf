terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }

    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.7.0"
    }

  }

  backend "s3" {}
}

provider "digitalocean" {
  token = var.digitalocean_token
}

provider "kubectl" {
  load_config_file       = false
  host                   = digitalocean_kubernetes_cluster.kube_cluster.endpoint
  token                  = digitalocean_kubernetes_cluster.kube_cluster.kube_config.0.token
  cluster_ca_certificate = base64decode(digitalocean_kubernetes_cluster.kube_cluster.kube_config.0.cluster_ca_certificate)
}

resource "digitalocean_kubernetes_cluster" "kube_cluster" {
  name    = var.cluster_name
  region  = "sgp1"
  version = var.cluster_version

  node_pool {
    name       = var.node_pool_name
    size       = var.node_pool_size
    node_count = var.node_pool_node_count
  }
}

data "kubectl_filename_list" "manifests" {
  pattern = "./manifests/*.yaml"
}

resource "kubectl_manifest" "kube_manifest" {
  count     = length(data.kubectl_filename_list.manifests.matches)
  yaml_body = file(element(data.kubectl_filename_list.manifests.matches, count.index))
}
