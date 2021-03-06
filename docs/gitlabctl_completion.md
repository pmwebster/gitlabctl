## gitlabctl completion

Generates bash completion scripts

### Synopsis

To configure your BASH shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(gitlabctl completion --bash)

To configure your ZSH shell to load completions for each session do the following

gitlabctl completion --zsh > /usr/local/share/zsh/site-functions/_gitlabctl

```
gitlabctl completion [flags]
```

### Options

```
      --bash   Generate bash completion script
  -h, --help   help for completion
      --zsh    Generate zsh completion script
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.gitlabctl.yaml)
```

### SEE ALSO

* [gitlabctl](gitlabctl.md)	 - Gitlab command-line interface

