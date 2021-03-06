terraform {
  required_providers {
    cheesecake = {
      versions = ["0.2"]
      source   = "example.com/joerx/cheesecake"
    }
  }
}

provider "cheesecake" {}

module "cake" {
  source = "./cheesecake"
}

output "cheesecakes" {
  value = module.cake.cheesecakes
}
