## gitlabctl get users

List all Gitlab users

### Synopsis

List all Gitlab users

```
gitlabctl get users [flags]
```

### Examples

```
# get all users
gitlabctl get users --out json

# get all users created after a date
gitlabctl get users --created-after="06/04/2018 22:05"


```

### Options

```
      --active                  Lookup users with active status (default true)
      --blocked                 Lookup users with blocked status (default true)
      --created-after string    Lookup users that are created after the specified value.
                                Example: gitlabctl get users --created-after=2001-01-02T00:00:00.060Z
                                Acceptable Format Reference: https://github.com/devopsctl/gitlabctl/blob/master/docs/dateparse.md
                                
      --created-before string   Lookup users that are created before the specified value.
                                Example: gitlabctl get users --created-before=2001-01-02T00:00:00.060Z
                                Acceptable Format Reference: https://github.com/devopsctl/gitlabctl/blob/master/docs/dateparse.md
                                
      --external-uid string     Lookup users by external uid.Best combined with --provider flag.
  -h, --help                    help for users
      --order-by string         Return projects ordered by id, name, username, created_at, updated_at  fields. Default is created_at (default "id")
      --provider string         Lookup users by provider. Best combined with --external-uid flag.
      --search string           Return the list of resources matching the search criteria
      --sort string             Order resources in asc or desc order. Default is asc (default "asc")
      --username string         Lookup users by username
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

