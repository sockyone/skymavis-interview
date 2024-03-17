locals {
  team_data = yamldecode(file("${path.module}/team.yaml"))
}

module "github_team" {
  source = "./modules/github_team"
  for_each = local.team_data
  team_name = each.key
  team_description = each.value.description
  members = each.value.members
  maintainers = each.value.maintainers

}