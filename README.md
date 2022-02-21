# accio: Go dependency manager

━☆ﾟ

# SUMMARY

accio extends Go to track buildtime dependencies.

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

accio automates away the low level `go install` commands involved in provisioning reproducible Go build environments.

Notably, accio's own runtime stack is small, requiring no other programming languages or scripts to operate.

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* any version control clients necessary for your development dependency tree (e.g., [git](https://git-scm.com/))

# DOCUMENTATION

https://godoc.org/github.com/mcandre/accio

# DOWNLOAD

https://github.com/mcandre/accio/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/accio/cmd/accio@latest
```

# UNINSTALL

```console
$ rm "$GOPATH/bin/accio"
```

# CONTRIBUTING

For more information on developing accio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# DETAIL

Some tools have not yet updated to use `go mod` (Go v1.11+ modules) for library dependency management. You can customize the package `go111module`, e.g. `go111module: "off"`, which activates the corresponding `GO111MODULE` environment variable configuration for the package.

Go does not support version pins for pre-Go v1.11 module packages.
