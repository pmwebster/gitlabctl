## gitlabctl new branch

Create a new branch for a specified project

### Synopsis

Create a new branch for a specified project

```
gitlabctl new branch [flags]
```

### Examples

```
# create a develop branch from master branch for project groupx/myapp
gitlabctl new branch develop --project=groupx/myapp --ref=master
```

### Options

```
  -h, --help             help for branch
  -p, --project string   The name or ID of the project
  -r, --ref string       The branch name or commit SHA to create branch from
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create a Gitlab resource

