## gitlabctl new group

Create a new group

### Synopsis

Create a new group

```
gitlabctl new group [flags]
```

### Options

```
  -h, --help                     help for group
      --lfs-enabled              enable LFS
      --name string              resource name
      --namespace string         The parent group id or group path if creating a subgroup. (defaults to current user namespace)
      --request-access-enabled   enable request access
      --visibility string        public, internal or private (default "private")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
  -o, --out string      Print the command output to the desired format. (json, yaml, simple) (default "simple")
```

### SEE ALSO

* [gitlabctl new](gitlabctl_new.md)	 - Create Gitlab resources

###### Auto generated by spf13/cobra on 26-May-2018