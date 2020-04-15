provider "aws" {
    region = "eu-west-1"
}

provider "digitalocean" {
    token = "${var.do_token}"
}

resource "digitalocean_database_cluster" "tutorialedge-db" {
    name       = "tutorialedge-db"
    engine     = "mysql"
    version    = "8"
    size       = "db-s-1vcpu-1gb"
    region     = "nyc1"
    node_count = 1
}