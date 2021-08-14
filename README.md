# accio: a dependency manager for Go developer tools

━☆ﾟ

# EXAMPLE

```console
$ cd example

$ which shadow
shadow not found

$ accio -install

$ which shadow
/home/andrew/go/bin/shadow
```

accio processes packages recursively.

See `accio -help` for more options.

## Configuration

For more detail on managing Go development packages, see [CONFIGURATION.md](CONFIGURATION.md).

# ABOUT

Buildtime dependency tools like golint, Mage, shadow, and so on involve executable artifacts, which `go mod` unfortunately silently ignores. Fortunately, we have accio to manage these kinds of dependencies.

accio automates away the low level `go get` / `git checkout` commands involved in provisioning reproducible Go build environments.

Notably, accio's own runtime stack is small, requiring no other programming languages or scripts to operate.

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

* [Go](https://golang.org/) 1.13+
* [git](https://git-scm.com/) 2.25+

# DOCUMENTATION

https://godoc.org/github.com/mcandre/accio

# INSTALL FROM SOURCE

```console
$ GO111MODULE=off go get github.com/mcandre/accio/cmd/accio
```

# UNINSTALL

```console
$ rm "$GOPATH/bin/accio"
```

# LIMITATIONS

* `version` overrides are implemented only for dependencies with git VCS repositories available.

# CONTRIBUTING

For more information on developing accio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).
