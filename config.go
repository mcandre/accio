package accio

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"gopkg.in/yaml.v2"
)

// ConfigFilename specifies the basic location of configuration data.
const ConfigFilename = "accio.yaml"

// Config controls accio behavior.
type Config struct {
	// Debug toggles logging.
	Debug bool

	// Packages collects development dependencies.
	Packages []Package

	// goPath caches the GOPATH environment variable. (Required)
	goPath string
}

// Load reads configuration data into memory.
func Load() (*Config, error) {
	goPath := os.Getenv("GOPATH")

	if goPath == "" {
		return nil, errors.New("missing environment variable GOPATH")
	}

	var config Config

	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	pth := path.Join(cwd, ConfigFilename)

	reader, err := os.Open(pth)

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(reader)

	if err2 := decoder.Decode(&config); err2 != nil {
		return nil, err2
	}

	config.goPath = goPath

	for i, p := range config.Packages {
		if p.Get == "" {
			return nil, errors.New("get cannot be blank")
		}

		if p.Executables == nil {
			p.Executables = &[]string{
				path.Base(p.Get),
			}
		}

		p.destination = path.Join(config.goPath, "src", p.Get)
		config.Packages[i] = p
	}

	return &config, nil
}

// CheckoutVersion applies the configured version onto the local clone for a given package.
func (o Config) CheckoutVersion(pkg Package) error {
	cmd := exec.Command("git")
	cmd.Args = []string{"git", "fetch", "--all"}
	cmd.Env = os.Environ()
	cmd.Dir = pkg.destination
	cmd.Stderr = os.Stderr

	if o.Debug {
		cmd.Stdout = os.Stdout
	}

	if o.Debug {
		log.Printf("Command: %v\n", cmd)
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git")
	cmd.Args = []string{"git", "checkout", "-f", "--detach", pkg.Version}
	cmd.Env = os.Environ()
	cmd.Dir = pkg.destination
	cmd.Stderr = os.Stderr

	if o.Debug {
		cmd.Stdout = os.Stdout
	}

	if o.Debug {
		log.Printf("Command: %v\n", cmd)
	}

	return cmd.Run()
}

// GoInstallPackage builds and places a package using `go install`.
func (o Config) GoInstallPackage(pkg Package) error {
	cmd := exec.Command("go")
	cmd.Args = []string{"go", "install", "./..."}
	cmd.Env = os.Environ()
	cmd.Dir = pkg.destination
	cmd.Stderr = os.Stderr

	if o.Debug {
		cmd.Stdout = os.Stdout
	}

	if o.Debug {
		log.Printf("Command: %v\n", cmd)
	}

	return cmd.Run()
}

// DownloadViaGoSource downloads the source for a given package using `go get`.
func (o Config) DownloadViaGoSource(pkg Package) error {
	args := []string{
		"go",
		"get",
		"-d",
	}

	if pkg.Update {
		args = append(args, "-u")
	}

	args = append(args, fmt.Sprintf("%s/...", pkg.Get))

	cmd := exec.Command("go")
	cmd.Args = args
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GO111MODULE=off")
	cmd.Stderr = os.Stderr

	if o.Debug {
		cmd.Stdout = os.Stdout
	}

	if o.Debug {
		log.Printf("Command: %v\n", cmd)
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	if pkg.Version == "" {
		return nil
	}

	return o.CheckoutVersion(pkg)
}

// DownloadViaGitSource acquires the source for a given package using `git`.
func (o Config) DownloadViaGitSource(pkg Package) error {
	args := []string{
		"git",
		"clone",
	}

	if pkg.Version != "" {
		args = append(args, "-b")
		args = append(args, pkg.Version)
	}

	args = append(args, pkg.URL)
	args = append(args, pkg.destination)

	cmd := exec.Command("git")
	cmd.Args = args
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr

	if o.Debug {
		cmd.Stdout = os.Stdout
	}

	if o.Debug {
		log.Printf("Command: %v\n", cmd)
	}

	return cmd.Run()
}

// InstallConfig acquires a toolset.
func (o Config) InstallPackage(pkg Package) error {
	if pkg.URL == "" {
		if err := o.DownloadViaGoSource(pkg); err != nil {
			return err
		}
	} else if err := o.DownloadViaGitSource(pkg); err != nil {
		return err
	}

	return o.GoInstallPackage(pkg)
}

// Install acquires the configured Go tools.
func (o Config) Install() error {
	for _, pkg := range o.Packages {
		if err := o.InstallPackage(pkg); err != nil {
			return err
		}
	}

	cmd := exec.Command("go")
	cmd.Args = []string{"go", "mod", "tidy"}
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr

	if o.Debug {
		cmd.Stdout = os.Stdout
	}

	if o.Debug {
		log.Printf("Command: %v\n", cmd)
	}

	return cmd.Run()
}

// Remove ensures the specified Go executable is not present in $GOPATH/bin.
func (o Config) Remove(executable string) error {
	pth := path.Join(o.goPath, "bin", Delve(executable))

	_, err := os.Stat(pth)

	if os.IsNotExist(err) {
		return nil
	}

	return os.Remove(pth)
}

// Destructo deletes all the configured Go tools.
func (o Config) Destructo() error {
	for _, pkg := range o.Packages {
		for _, executable := range *pkg.Executables {
			if err := o.Remove(executable); err != nil {
				return err
			}
		}
	}

	return nil
}
