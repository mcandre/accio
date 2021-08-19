# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* a POSIX compatible shell (e.g., `bash`, `ksh`, `sh`, `zsh`)
* Go development tools (`sh acquire`)
* [zip](https://linux.die.net/man/1/zip)

# INSTALL FROM SOURCE

```console
$ go install ./...
```

# TEST

```console
$ cd example
$ accio -install
```

# PORT

```console
$ FACTORIO_BANNER=accio-0.0.1 factorio

$ cd bin

$ zip -r accio-0.0.1.zip accio-0.0.1
```
