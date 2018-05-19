![](https://talks.golang.org/2012/waza/gophercomplex1.jpg)
# Gitlab Control (Work in Progress)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopsctl/gitlabctl)](https://goreportcard.com/report/github.com/devopsctl/gitlabctl)

<!-- vim-markdown-toc GFM -->

* [Custom Packages](#custom-packages)
* [Test Driven Development](#test-driven-development)
* [How the Commands Authenticate](#how-the-commands-authenticate)
* [Group Command - `group`](#group-command---group)
	* [List all groups - `ls`](#list-all-groups---ls)
	* [List all the subgroups of a group - `ls-subgroup`](#list-all-the-subgroups-of-a-group---ls-subgroup)
	* [List all the projects of a group - `ls-project`](#list-all-the-projects-of-a-group---ls-project)
	* [Delete a group - `rm`](#delete-a-group---rm)
	* [Create a group - `new`](#create-a-group---new)
	* [Edit a group - `edit`](#edit-a-group---edit)
	* [List all members of a group - `ls-member`](#list-all-members-of-a-group---ls-member)
	* [Remove a group member - `rm-member`](#remove-a-group-member---rm-member)
	* [Add a group member - `new-member`](#add-a-group-member---new-member)
	* [Remove all group members - `rm-all-member`](#remove-all-group-members---rm-all-member)
* [Project Service Command - `project`](#project-service-command---project)
	* [List all projects - `ls`](#list-all-projects---ls)
	* [Delete a project - `rm`](#delete-a-project---rm)
	* [Create a project - `new`](#create-a-project---new)
	* [Edit a project - `edit`](#edit-a-project---edit)
	* [List all members of a project - `ls-member`](#list-all-members-of-a-project---ls-member)
	* [Remove a project member - `rm-member`](#remove-a-project-member---rm-member)
	* [Add a project member - `new-member`](#add-a-project-member---new-member)
	* [Remove all project members - `rm-all-member`](#remove-all-project-members---rm-all-member)
	* [List all hooks of a project - `ls-hooks`](#list-all-hooks-of-a-project---ls-hooks)
	* [Add a project hook - `new-hook`](#add-a-project-hook---new-hook)
	* [Edit a project hook - `edit-hook`](#edit-a-project-hook---edit-hook)
	* [Delete a project hook - `rm-hook`](#delete-a-project-hook---rm-hook)
	* [Delete all hooks in a project - `rm-all-hook`](#delete-all-hooks-in-a-project---rm-all-hook)

<!-- vim-markdown-toc -->

Our goal is to create a gitlab cli written in Go that is simple to use and easy to maintain. The code must be simple and flags must be patterned with the gitlab client package https://godoc.org/github.com/xanzy/go-gitlab.

It is worth noting that there are existing gitlab cli written in Go.

* https://github.com/michaellihs/golab (most complete but code is complex - at least for me to understand..)
* https://github.com/kyokomi/gitlab-cli (incomplete)
* https://github.com/clns/gitlab-cli (incomplete)

## Custom Packages

* Gitlab api client - https://godoc.org/github.com/xanzy/go-gitlab 
* Commandline flags - https://github.com/spf13/cobra 

## Test Driven Development

This project may grow big in the future so the definition of done for every commands should be tested against a local gitlab instance. 

## How the Commands Authenticate

Authenticate using environment variables.

* Basic authentication - `GITLAB_USERNAME`, `GITLAB_PASSWORD` and `GITLAB_HTTP_URL`
* Private token authentication - `GITLAB_PRIVATE_TOKEN` and `GITLAB_API_HTTP_URL`
* OAuth2 token authentication - `GITLAB_OAUTH_TOKEN` and `GITLAB_API_HTTP_URL`

## Group Command - `group`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupService

* [List all groups - `ls`](#list-all-groups---ls)
* [List group's subgroup - `ls-subgroup`](#list-groups-subgroup---ls-subgroup)
* [List group's projects - `ls-project`](#list-groups-projects---ls-projects)
* [Delete a group - `rm`](#delete-a-group---delete)
* [Create a group - `new`](#create-a-group---create)
* [Edit a group - `edit`](#edit-a-group---edit)

### List all groups - `ls`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListGroups

This sub command does not require a flag.

### List all the subgroups of a group - `ls-subgroup`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListSubgroups

| Flag  | Type   | Description                                                    | Required? | Default |
| :---- | :---   | :----------                                                    | :-------- | :------ |
| path  | string | the group name, id or full the path including the parent group | yes       |

### List all the projects of a group - `ls-project`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListGroupProjects

| Flag  | Type   | Description                                                    | Required? | Default |
| :---- | :---   | :----------                                                    | :-------- | :------ |
| path  | string | the group name, id or full the path including the parent group | yes       |

### Delete a group - `rm`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.DeleteGroup

| Flag  | Type   | Description                                                    | Required? |
| :---- | :---   | :----------                                                    | :-------- |
| path  | string | the group name, id or full the path including the parent group | yes       |

### Create a group - `new`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.CreateGroup

| Flag                   | Type   | Description                                               | Required? | Default |
| :----                  | :---   | :----------                                               | :-------- | :------ |
| name                   | string | The group name                                            | yes       |         |
| namespace              | string | The parent group id or group path if creating a subgroup. | no        |         |
| visibility             | string | public, internal or private                               | no        |
| lfs-enabled            | bool   | Enable LFS                                                | no        |
| request-access-enabled | bool   | Enable Request Access                                     | no        |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.

### Edit a group - `edit`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.UpdateGroup

| Flag                   | Type   | Description                   | Required? | Default |
| :----                  | :---   | :----------                   | :-------- | :------ |
| path  | string | the group name, id or full the path including the parent group | yes       |
| visibility             | string | public, internal or private   | no        |
| lfs-enabled            | bool   | Enable LFS                    | no        |
| request-access-enabled | bool   | Enable Request Access         | no        |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.
* Command requires at least 1 optional flag to be set.

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupMembersService

### List all members of a group - `ls-member` 

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupsService.ListGroupMembers

| Flag  | Type   | Description                                                    | Required? | Default |
| :---- | :---   | :----------                                                    | :-------- | :------ |
| path  | string | the group name, id or full the path including the parent group | yes       |         |

### Remove a group member - `rm-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupMembersService.RemoveGroupMember

| Flag     | Type   | Description                                                    | Required? | Default |
| :----    | :---   | :----------                                                    | :-------- | :------ |
| path     | string | the group name, id or full the path including the parent group | yes       |         |
| username | string | username to remove                                             | yes       |         |

### Add a group member - `new-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#GroupMembersService.AddGroupMember

| Flag         | Type   | Description                                                                                                           | Required? | Default |
| :----        | :---   | :----------                                                                                                           | :-------- | :------ |
| path         | string | the group name, id or full the path including the parent group                                                        | yes       |         |
| username     | string | username to add                                                                                                       | yes       |         |
| access-level | int    | member group access level (0, 10, 20, 30, 40, 50). Reference: https://docs.gitlab.com/ce/permissions/permissions.html | no        | 10      |

### Remove all group members - `rm-all-member`

A wrapper of listing all group members and deleting them all.

| Flag         | Type   | Description                                                                                                           | Required? | Default |
| :----        | :---   | :----------                                                                                                           | :-------- | :------ |
| path         | string | the group name, id or full the path including the parent group                                                        | yes       |         |
| username     | string | username to add                                                                                                       | yes       |         |

## Project Service Command - `project`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService

| Flag  | Type | Description | Required? | Default |
| :---- | :--- | :---------- | :-------- | :------ |

### List all projects - `ls`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.ListProjects

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Delete a project - `rm`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.DeleteProject

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Create a project - `new`

| Flag                                        | Type   | Description                                                                                     | Required? | Default |
| :----                                       | :---   | :----------                                                                                     | :-------- | :------ |
| name                                        | string | The project name                                                                                | yes       |         |
| namespace                                   | string | The parent group id or group path if creating a subgroup.                                       | no        |         |
| description                                 | string | project description                                                                             | no        |         |
| issues-enabled                              | bool   | enable issues                                                                                   | no        |
| merge-requests-enabled                      | bool   | enable merge requests                                                                           | no        |
| jobs-enabled                                | bool   | enable jobs                                                                                     | no        |
| wiki-enabled                                | bool   | enable wikis                                                                                    | no        |
| snippets-enabled                            | bool   | enable snippets                                                                                 | no        |
| resolve-outdated-diff-discussions           | bool   | resolve outdated diff discussions                                                               | no        |
| container-registry-enabled                  | bool   | enable container registry                                                                       | no        |
| shared-runners-enabled                      | bool   | enable shared runners                                                                           | no        |
| visibility                                  | string | project visibility (public, internal, private)                                                  | no        | public  |
| public-jobs                                 | bool   | if true, jobs can be viewed by non-project-members                                              | no        |
| only-allow-merge-if-pipeline-succeeds       | bool   | set whether merge requests can only be merged with successful jobs                              | no        |
| only-allow-merge-if-discussion-are-resolved | bool   | set whether merge requests can only be merged when all the discussions are resolved             | no        |
| merge-method                                | string | set the merge method used                                                                       | no        |
| lfs-enabled                                 | bool   | enable lfs                                                                                      | no        |
| request-access-enabled                      | bool   | allow users to request member access                                                            | no        |
| tag-list                                    | string | the list of tags for a project; put array of tags, that should be finally assigned to a project | no        |
| printing-merge-request-link-enabled         | bool   | show link to create/view merge request when pushing from the command line                       | no        |
| ci-config-path                              | string | the path to ci config file                                                                      | no        |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.

### Edit a project - `edit`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.EditProject

| Flag                                        | Type   | Description                                                                                     | Required? | Default |
| :----                                       | :---   | :----------                                                                                     | :-------- | :------ |
| path                                        | string | the project name, id or full the path including the parent group - (path/to/project)            | yes       |         |
| description                                 | string | project description                                                                             | no        |         |
| issues-enabled                              | bool   | enable issues                                                                                   | no        |
| merge-requests-enabled                      | bool   | enable merge requests                                                                           | no        |
| jobs-enabled                                | bool   | enable jobs                                                                                     | no        |
| wiki-enabled                                | bool   | enable wikis                                                                                    | no        |
| snippets-enabled                            | bool   | enable snippets                                                                                 | no        |
| resolve-outdated-diff-discussions           | bool   | resolve outdated diff discussions                                                               | no        |
| container-registry-enabled                  | bool   | enable container registry                                                                       | no        |
| shared-runners-enabled                      | bool   | enable shared runners                                                                           | no        |
| visibility                                  | string | project visibility (public, internal, private)                                                  | no        | public  |
| public-jobs                                 | bool   | if true, jobs can be viewed by non-project-members                                              | no        |
| only-allow-merge-if-pipeline-succeeds       | bool   | set whether merge requests can only be merged with successful jobs                              | no        |
| only-allow-merge-if-discussion-are-resolved | bool   | set whether merge requests can only be merged when all the discussions are resolved             | no        |
| merge-method                                | string | set the merge method used                                                                       | no        |
| lfs-enabled                                 | bool   | enable lfs                                                                                      | no        |
| request-access-enabled                      | bool   | allow users to request member access                                                            | no        |
| tag-list                                    | string | the list of tags for a project; put array of tags, that should be finally assigned to a project | no        |
| printing-merge-request-link-enabled         | bool   | show link to create/view merge request when pushing from the command line                       | no        |
| ci-config-path                              | string | the path to ci config file                                                                      | no        |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.
* Command requires at least 1 optional flag to be set.

### List all members of a project - `ls-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectMembersService.ListProjectMembers

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Remove a project member - `rm-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectMembersService.AddProjectMember

| Flag     | Type   | Description                                                                          | Required? | Default |
| :----    | :---   | :----------                                                                          | :-------- | :------ |
| path     | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| username | string | the new member username                                                              | yes       |         |

### Add a project member - `new-member`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectMembersService.AddProjectMember

| Flag         | Type   | Description                                                                                                           | Required? | Default |
| :----        | :---   | :----------                                                                                                           | :-------- | :------ |
| path         | string | the project name, id or full the path including the parent group - (path/to/project)                                  | yes       |         |
| username     | string | the new member username                                                                                               | yes       |         |
| access-level | int    | member group access level (0, 10, 20, 30, 40, 50). Reference: https://docs.gitlab.com/ce/permissions/permissions.html | no        | 10      |

### Remove all project members - `rm-all-member`

A wrapper of listing all project members and removing them all.

| Flag     | Type   | Description                                                                          | Required? | Default |
| :----    | :---   | :----------                                                                          | :-------- | :------ |
| path     | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### List all hooks of a project - `ls-hooks`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.ListProjectHooks

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Add a project hook - `new-hook`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.AddProjectHook

| Flag                       | Type   | Description                                                                          | Required? | Default |
| :----                      | :---   | :----------                                                                          | :-------- | :------ |
| path                       | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| url                        | string | The hook URL                                                                         | yes       |
| push-events                | bool   | Trigger hook on push events                                                          |           |
| issues-events              | bool   | Trigger hook on issues events                                                        |           |
| confidential-issues-events | bool   | Trigger hook on confidential issues events                                           |           |
| merge-requests-events      | bool   | Trigger hook on merge requests events                                                |           |
| tag-push-events            | bool   | Trigger hook on tag push events                                                      |           |
| note-events                | bool   | Trigger hook on note events                                                          |           |
| job-events                 | bool   | Trigger hook on wiki events                                                          |           |
| pipeline-events            | bool   | Trigger hook on pipeline events                                                      |           |
| wiki-page-events           | bool   | Trigger hook on wiki events                                                          |           |
| enable-ssl-verification    | bool   | Do SSL verification when triggering the hook                                         |           |
| token                      | string | Secret token to validate received payloads                                           |           |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.

### Edit a project hook - `edit-hook`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.EditProjectHook

| Flag                       | Type   | Description                                                                          | Required? | Default |
| :----                      | :---   | :----------                                                                          | :-------- | :------ |
| id                         | int    | hook id                                                                              | yes       |         |
| path                       | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |
| url                        | string | The hook URL                                                                         | yes       |
| push-events                | bool   | Trigger hook on push events                                                          |           |
| issues-events              | bool   | Trigger hook on issues events                                                        |           |
| confidential-issues-events | bool   | Trigger hook on confidential issues events                                           |           |
| merge-requests-events      | bool   | Trigger hook on merge requests events                                                |           |
| tag-push-events            | bool   | Trigger hook on tag push events                                                      |           |
| note-events                | bool   | Trigger hook on note events                                                          |           |
| job-events                 | bool   | Trigger hook on wiki events                                                          |           |
| pipeline-events            | bool   | Trigger hook on pipeline events                                                      |           |
| wiki-page-events           | bool   | Trigger hook on wiki events                                                          |           |
| enable-ssl-verification    | bool   | Do SSL verification when triggering the hook                                         |           |
| token                      | string | Secret token to validate received payloads                                           |           |

Custom flag validation:

* If optional or non-required flags are not set, do not use or ignore the default value.
* Command requires at least 1 optional flag to be set.

### Delete a project hook - `rm-hook`

API doc: https://godoc.org/github.com/xanzy/go-gitlab#ProjectsService.DeleteProjectHook

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| id    | int    | hook id                                                                              | yes       |         |
| path  | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |

### Delete all hooks in a project - `rm-all-hook`

A wrapper of listing all project hooks and deleting all of them.

| Flag  | Type   | Description                                                                          | Required? | Default |
| :---- | :---   | :----------                                                                          | :-------- | :------ |
| path  | string | the project name, id or full the path including the parent group - (path/to/project) | yes       |         |


