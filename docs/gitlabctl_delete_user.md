## gitlabctl delete user

Delete a Gitlab user by specifying the username

### Synopsis

Delete a Gitlab user by specifying the username

```
gitlabctl delete user [flags]
```

### Examples

```
# delete a user by username
gitlabctl delete user john.smith

# delete a user with user id (15)
gitlabctl delete user 15
```

### Options

```
  -h, --help   help for user
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl delete](gitlabctl_delete.md)	 - Delete a Gitlab resource

