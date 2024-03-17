variable team_name {
  type = string
}

variable "team_description" {
  type = string
}

variable "members" {
  type = list
  description = "list members in team"
}

variable "maintainers" {
  type = list
  description = "list maintainers in team"
}