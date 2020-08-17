terraform {
  required_providers {
    cheesecake = {
      versions = ["0.2"]
      source   = "example.com/joerx/cheesecake"
    }
  }
}

data "cheesecake_cheesecakes" "all" {}

# Returns all coffees
output "cheesecakes" {
  value = data.cheesecake_cheesecakes.all.cheesecakes
}
