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

# WARNINGS

accio overwrites the directories `$GOPATH/src/<your tool dependencies>` when provisioning your configured Go tools. If you contribute to your development tools, take care to regularly push your changes to a remote repository. There is an inherent risk that any local, unpushed changes may be overwritten by `accio -install`.

For this reason, it is recommended to not bootstrap your project's development by depending on itself in terms of tooling. Nor depend on any previous version of itself. Don't create a dependency cycle, which accio may corrupt. Likewise, try not to create an indirect dependency cycle between multiple tools.

Each dev tool should be simple to compile and install from scratch, using plain `go install [./...]` commands. Go packages which do not support the standard `go get`, `go mod`, `go generate`, etc. built-in commands for building and installing, may not integrate well (with accio or any other Go projects).

Go binaries can be helpful for prototyping development tools, but reusable Go libraries will naturally integrate best with `go mod`. For example, invoke development libraries as auxilliary build tasks using the [Mage](https://magefile.org/) task runner. Any development libraries you reference in your `magefile.go` will automatically be tracked with `go.mod` and will not need accio.

If your project requires vendoring the source code for the tools, then replace the shellouts with pure Go API calls. A Mage file can issue these API calls in a way that integrates more naturally with `go mod`. At which point you can remove that command line tool from your accio.yaml. Finally, invoke `go mod download; go mod tidy; go mod vendor`.

For some Go tools it can take too much effort to lib-ize a command line tool. In that case, establish an independent fork of the tool's primary repository, and override the package's URL field in your accio configuration to point to your fork.
