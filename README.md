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

# WARNINGS

accio overwrites the directories `$GOPATH/src/<your tool dependencies>` when provisioning your configured Go tools. If you contribute to your development tools, take care to regularly push your changes to a remote repository. There is an inherent risk that any local, unpushed changes may be overwritten by `accio -install`.

For this reason, it is recommended to not bootstrap your project's development by depending on itself in terms of tooling. Nor depend on any previous version of itself. Don't create a dependency cycle, which accio may corrupt. Likewise, try not to create an indirect dependency cycle between multiple tools.

Each dev tool should be simple to compile and install from scratch, using plain `go install [./...]` commands. Go packages which do not support the standard `go get` / `go mod` system for building and installing, may not work with accio.

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

* [Go](https://golang.org/) 1.13+
* any version control clients necessary for your development dependency tree (e.g., [git](https://git-scm.com/))

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

# CONTRIBUTING

For more information on developing accio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).
