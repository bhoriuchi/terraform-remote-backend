terraform {
  backend "remote" {
    hostname = "localhost:9443"
    organization = "blackline"

    workspaces {
      name = "workspaces/nonprod/test"
    }
  }
}

resource "local_file" "foo" {
  content = "foo"
  filename = "/tmp/foo.txt"
}