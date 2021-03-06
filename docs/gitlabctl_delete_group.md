## gitlabctl delete group

Delete a Gitlab group by specifying the id or group path

### Synopsis

Delete a Gitlab group by specifying the id or group path

```
gitlabctl delete group [flags]
```

### Examples

```
# delete a Group named GroupX
gitlabctl delete group GroupX

# delete a Subgroup named GroupY under GroupX
gitlabctl delete group GroupX/GroupY

# delete a group with id (3)
gitlabctl delete group 3
```

### Options

```
  -h, --help   help for group
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

