variable "gogs_token" {
    type = "string"
}

variable "gogs_base_url" {
    type = "string"
}

provider "gogs" {
    token = "${var.gogs_token}"
    base_url = "${var.gogs_base_url}"
}

resource "gogs_user" "bob" {
    login = "bob"
    password = "password"
    email = "bob@example.com"
    full_name = "Bob Marley"
}