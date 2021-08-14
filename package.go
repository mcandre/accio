package accio

// Package bundles salient details for Go dependencies, particularly development tools.
type Package struct {
	// Get denotes the path one uses to `go get` a toolset. (Required)
	//
	// Example: "golang.org/x/tools/go/analysis/passes/shadow"
	Get string

	// URL denotes a network path override. (Default: `go get` determined URL)
	//
	// Example: ""
	// Example: "https://github.com/golang/tools.git"
	URL string

	// Version denotes a specific version control reference, such as a tag, branch, or commit. (Default: Latest according to `go get` determined cache semantics)
	//
	// Example: ""
	// Example: "3.14.0"
	Version string

	// Update denotes whether to forcibly attempt to search for and apply updates with the -u flag to `go get`. (Default: false)
	//
	// Update is ignored when a Version is specified.
	//
	// Example: true
	// Example: false
	Update bool

	// Executables denotes the filenames for any executable artifacts, excluding file extensions. (Default: [Get tail])
	//
	// Example: nil
	// Example: ["shadow"]
	Executables *[]string

	// destination is populated by $GOPATH/src/<Get>.
	destination string
}
