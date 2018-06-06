[![Throughput Graph](http://graphs.waffle.io/devopsctl/gitlabctl/throughput.svg)](https://waffle.io/devopsctl/gitlabctl/metrics)

[![Build Status](https://travis-ci.org/devopsctl/gitlabctl.svg?branch=master)](https://travis-ci.org/devopsctl/gitlabctl)
[![codecov](https://codecov.io/gh/devopsctl/gitlabctl/branch/master/graph/badge.svg)](https://codecov.io/gh/devopsctl/gitlabctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopsctl/gitlabctl)](https://goreportcard.com/report/github.com/devopsctl/gitlabctl)
[![Waffle board](https://badge.waffle.io/devopsctl/gitlabctl.png?columns=all)](https://waffle.io/devopsctl/gitlabctl)

# gitlabctl - work in progress

A Gitlab command-line interface.

## Getting Started

Please see [Github site](https://devopsctl.github.io/gitlabctl/) for a nice documentation of this project.

## Installing

Download the binary from [releases](https://github.com/devopsctl/gitlabctl/releases) or use Go Get, `go get github.com/devopsctl/gitlabctl`.

## Contributing

Contributors are welcomed with love! Please read [CONTRIBUTING.md](./CONTRIBUTION.md) for the process for submitting pull requests to us.

## Gitlab Commands Available

### Authentication 

* [x] Authentication through environment variables. See [gitlabctl](gitlabctl.md).
* [ ] Authentication using `gitlabctl login` command.

### Group

* [x] `get groups [flags]`
* [x] `get groups --from-group [flags]`
* [x] `describe group [group path] [flags]`
* [x] `new group [group name] [flags]`
* [x] `delete group [group path]`
* [x] `edit group [group  name] [flags]`

### Project

* [x] `get projects [flags]`
* [x] `get projects --from-group  [flags]`
* [x] `describe project [project path]`
* [x] `new project [project name] [flags]`
* [x] `edit project [project name] [flags]`
* [x] `delete project [project path]`

### Project Hooks

* [ ] `get project-hooks [project name]`
* [ ] `delete all-project-hook [project name]`
* [ ] `delete project-hook [hook id]`
* [ ] `new project-hook [flags]`
* [ ] `edit project-hook [flags]`

### Users 

* [x] `get users [flags]`
* [x] `describe user [username]`
* [x] `new user [username] [flags]`
* [x] `delete user [username]`
* [ ] `edit user [username] [flags]`

### Users SSH Keys 

* [ ] `get ssh-keys`
* [ ] `get ssh-keys --from-user=[username]`
* [ ] `new ssh-key [flags]`
* [ ] `delete ssh-key [flags]`

### Users GPG Keys

* [ ] `get gpg-keys`
* [ ] `get gpg-keys --from-user=[username]`
* [ ] `new gpg-key [flags]`
* [ ] `delete gpg-key [flags]`

### Members

* [ ] `get members --from-group [flags]`
* [ ] `get members --from-project [flags]`
* [ ] `describe member [username] --from-group [flags]`
* [ ] `describe member [username] --from-project [flags]`
* [ ] `new member [username] --from-group [flags]`
* [ ] `new member [username] --from-project [flags]`
* [ ] `remove member [username] --from-group`
* [ ] `remove member [username] --from-project`
* [ ] `edit member [username] --from-group [flags]`
* [ ] `edit member [username] --from-project [flags]`

