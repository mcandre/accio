//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/accio"
	"github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Lint

// Audit runs a security audit.
func Audit() error { return mageextras.SnykTest() }

// GoVet runs go vet with shadow checks enabled.
func GoVet() error { return mageextras.GoVetShadow() }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Staticcheck() }

// Yamllint runs yamllint.
func Yamllint() error { return mageextras.Yamllint("-s", ".yamllint", ".") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Staticcheck)
	mg.Deps(Yamllint)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("accio-%s", accio.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/accio"

// Factorio cross-compiles Go binaries for a multitude of platforms.
func Factorio() error { return mageextras.Factorio(portBasename) }

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Factorio); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("accio") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
