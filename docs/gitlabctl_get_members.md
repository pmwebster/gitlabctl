## gitlabctl get members

List all members of a group or a project

### Synopsis

List all members of a group or a project

```
gitlabctl get members [flags]
```

### Examples

```
# list all members of a groups
gitlabctl get members --from-group Group1

# list all members of a project
gitlabctl get members --from-project Group1/Project1
```

### Options

```
  -G, --from-group string     Use a group as the target namespace when performing the command
  -P, --from-project string   Use a project as the target namespace when performing the command
  -h, --help                  help for members
  -q, --query string          A query string to search for members(defaults to blank)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
      --page int        Page of results to retrieve
      --per-page int    The number of results to include per page
```

### SEE ALSO

* [gitlabctl get](gitlabctl_get.md)	 - Get Gitlab resources

