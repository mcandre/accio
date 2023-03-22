package accio

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"gopkg.in/yaml.v3"
)

// ConfigFilename specifies the basic location of configuration data.
const ConfigFilename = "accio.yaml"

// Config controls accio behavior.
type Config struct {
	// Debug toggles logging.
	Debug bool `json:"debug" yaml:"debug"`

	// Packages collects development dependencies.
	Packages []Package `json:"packages" yaml:"packages"`

	// goPath caches the GOPATH environment variable. (Required)
	goPath string `json:"-" yaml:"-"`
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
		if p.Name == "" {
			return nil, errors.New("name cannot be blank")
		}

		if p.Executables == nil {
			p.Executables = &[]string{
				path.Base(p.Name),
			}
		}

		config.Packages[i] = p
	}

	return &config, nil
}

// InstallPackage acquires a toolset.
func (o Config) InstallPackage(pkg Package) error {
	pin := pkg.Name

	if pkg.Version != "" {
		pin = fmt.Sprintf("%s@%s", pkg.Name, pkg.Version)
	}

	cmd := exec.Command("go")
	cmd.Args = []string{"go", "install", pin}
	cmd.Env = os.Environ()

	if pkg.Go111Module != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("GO111MODULE=%s", pkg.Go111Module))
	}

	cmd.Stderr = os.Stderr

	if o.Debug {
		cmd.Stdout = os.Stdout
		log.Printf("command: %v\n", cmd)
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
