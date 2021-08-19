package accio

// Package bundles salient details for Go dependencies, particularly development tools.
type Package struct {
	// Name denotes the path one uses to `go install` a toolset. (Required)
	//
	// Example: "golang.org/x/tools/go/analysis/passes/shadow"
	// Example: "github.com/mcandre/stank/..."
	Name string `json:"name" yaml:"name"`

	// Go111Module propagates a GO111MODULE environment variable. (Optional)
	//
	// Example: ""
	// Example: "auto"
	// Example: "off"
	// Example: "on"
	Go111Module string `json:"go111module" yaml:"go111module"`

	// Version denotes a specific version control reference, such as a tag, branch, or commit. (Default: blank)
	//
	//
	// Go does not support version pins for pre-Go v1.11 module packages.
	//
	// Example: ""
	// Example: "latest"
	// Example: "v1.2.3"
	Version string `json:"version" yaml:"version"`

	// Executables denotes the filenames for any executable artifacts, excluding file extensions. (Default: [Install tail])
	//
	// Example: nil
	// Example: ["shadow"]
	Executables *[]string `json:"executables" yaml:"executables"`
}
