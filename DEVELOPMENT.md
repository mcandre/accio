# BUILDTIME REQUIREMENTS

* an [sh](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/sh.html) implementation
* [Go](https://golang.org/) 1.20.2+ with `sh acquire`
* [Node.js](https://nodejs.org/en) 16.14.2+ with `npm install -g snyk@1.996.0`
* [Python](https://www.python.org/) 3.11.2+ with `pip[3] install --upgrade pip setuptools` and `pip[3] install -r requirements-dev.txt`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10
* [direnv](https://direnv.net/) 2

# AUDIT

```console
$ mage audit
```

# INSTALL

```console
$ mage install
```

# UNINSTALL

```console
$ mage uninstall
```

# TEST

```console
$ cd example
$ accio -install
```

# PORT

```console
$ mage port
```
