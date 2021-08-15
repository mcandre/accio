package accio

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/Masterminds/vcs"
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

// InstallConfig acquires a toolset.
func (o Config) InstallPackage(pkg Package) error {
	if o.Debug {
		log.Printf("Installing package %s\n", pkg.Get)
	}

	_, err := os.Stat(pkg.destination)

	noSuchDirectory := os.IsNotExist(err)

	var remote string

	if noSuchDirectory {
		if pkg.URL != "" {
			remote = pkg.URL
		} else {
			remote = fmt.Sprintf("https://%s", pkg.Get)
		}
	}

	destination := path.Clean(pkg.destination)

	var repo vcs.Repo

	for {
		repo, err = vcs.NewRepo(remote, destination)

		if err == nil {
			break
		}

		if err == vcs.ErrCannotDetectVCS {
			if path.Dir(destination) == destination {
				break
			}

			destination = path.Dir(destination)
		}
	}

	if err != nil {
		return err
	}

	if noSuchDirectory {
		if err2 := repo.Get(); err2 != nil {
			return err2
		}
	}

	if pkg.Version != "" {
		if err2 := repo.UpdateVersion(pkg.Version); err2 != nil {
			return err2
		}
	}

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
