# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.17+
* a POSIX compatible shell (e.g., `bash`, `ksh`, `sh`, `zsh`)
* Go development tools (`sh acquire`)

## Recommended

* [snyk](https://www.npmjs.com/package/snyk) 1.893.0 (`npm install -g snyk@1.893.0`)
* [zip](https://linux.die.net/man/1/zip)

# SECURITY AUDIT

```console
$ snyk test
```

# INSTALL

```console
$ go install ./...
```

# UNINSTALL

```console
$ rm "$GOPATH/bin/accio"
```

# TEST

```console
$ cd example
$ accio -install
```

# PORT

```console
$ FACTORIO_BANNER=accio-0.0.2 factorio

$ cd bin

$ zip -r accio-0.0.2.zip accio-0.0.2
```
