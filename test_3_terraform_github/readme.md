# Manage Github teams

## Run
Role configuration will be placed in `team.yaml` file.
```
terraform init
terraform plan
```

## My thought about the problem
I think there are many ways to manage Github roles and teams. You may want to control everything from the view of team, or from the view of user:

- Team-oriented: Managed by teams. Want to assign user to a team? -> Go find team, add user to the team.
- User-oriented: Managed by users. Want to assign user to a team? -> Find user, add team name and role to the user.
- Flatten-view: There are just lines of records `(member, team, role)`.

In my organization, I saw people prefer managing in Team-oriented way. Sometime you want to check the permission of a particular user (for security audit or offboarding process maybe), but with my solution, you still can easily `ctrl F` and everything will show up.

> [!NOTE]
> If you have more teams, you may want to split the `team.yaml` file into many other files. Each team will manage their own file, and you can set the permission for editting these files with Gitlab file owner feature. For example: Leaders will have permission to modify files of their own team. You will merge these yaml files in the pipeline then pass it to terraform.
