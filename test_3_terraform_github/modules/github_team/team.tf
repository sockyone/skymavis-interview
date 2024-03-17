resource "github_team" "team" {
  name        = var.team_name
  description = var.team_description
  privacy     = "closed"
}

resource "github_team_membership" "role_members" {
  for_each = toset(var.members)

  team_id  = github_team.team.id
  username = each.key
  role     = "member"
}

resource "github_team_membership" "role_maintainers" {
  for_each = toset(var.maintainers)

  team_id  = github_team.team.id
  username = each.key
  role     = "maintainer"
}
