## gitlabctl

Gitlab command-line interface

### Synopsis

gitlabctl is a Gitlab client for the command-line.

This client helps you view, update, create, and delete Gitlab resources from the 
command-line interface.

There are two options to authenticate the command-line client to Gitlab interface:

1.) Using the 'login' command by passing the host url, username and password.

$ gitlabctl login

The login token will be saved in $HOME/.gitlabctl.yaml file.

2.) Using Environment variables.

* Basic Authentication (if using a username and password)
    - GITLAB_USERNAME
    - GITLAB_PASSWORD
    - GITLAB_HTTP_URL

* Private Token (if using a private token)
    - GITLAB_PRIVATE_TOKEN
    - GITLAB_API_HTTP_URL

* OAuth Token (if using an oauth token)
    - GITLAB_OAUTH_TOKEN
    - GITLAB_API_HTTP_URL

### Options

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -h, --help            help for gitlabctl
```

### SEE ALSO

* [gitlabctl completion](gitlabctl_completion.md)	 - Generates bash completion scripts
* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource
* [gitlabctl describe](gitlabctl_describe.md)	 - Describe a gitlab resource
* [gitlabctl edit](gitlabctl_edit.md)	 - Update or patch a Gitlab resource
* [gitlabctl gendoc](gitlabctl_gendoc.md)	 - Generate markdown documentation
* [gitlabctl get](gitlabctl_get.md)	 - Get Gitlab resources
* [gitlabctl login](gitlabctl_login.md)	 - Login to gitlab
* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource
* [gitlabctl version](gitlabctl_version.md)	 - 

