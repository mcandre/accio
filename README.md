# accio: Go dependency manager

# DEPRECATED

Recommend provisioning Go dev tools, or most any programming language dev tools for that matter, using a simple `all` task in a conventional `makefile`.

Be sure to followup your `go install`... commands with `go mod tidy` (and `modvendor -copy="**/*.h **/*.c **/*.hpp **/*.cpp"` in certain Cgo projects).

Try to write only dirt simple `go install`..., `pip3 install`..., `cargo install`... commands in the `makefile`. That way, your dev provisioning `makefile` script is more likely to succeed on other platforms.

Optionally, lint the makefile with `unmake` to detect more portability concerns.

https://github.com/mcandre/unmake

Or, if you happen to be using Go dev tools for non-Go projects, then you may already have a provisioning system available. For example, shell script projects can provision Go dev tools in a `./script`. (You may want to lint these with bashate, ShellCheck, shfmt, and stank.)

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

* [Go](https://golang.org/) 1.20.2+
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

# SEE ALSO

* [Ansible](https://www.ansible.com/)
* [Batsh](https://batsh.org/)
* [Docker](https://www.docker.com/)
* [tools.go](https://marcofranssen.nl/manage-go-tools-via-go-modules) (fragile)
