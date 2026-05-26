package ignore

// vcsDirectoryNames maps directory names to a boolean indicating whether or not
// they represent a VCS directory.
var vcsDirectoryNames = map[string]bool{
	".git":   true,
	".svn":   true,
	".hg":    true,
	".bzr":   true,
	"_darcs": true,
}

// vcsIgnorer is a wrapper Ignorer that provides VCS ignoring behavior.
type vcsIgnorer struct {
	// ignorer is the underlying ignorer.
	ignorer Ignorer
}

// Ignore implements Ignorer.Ignore.
func (i *vcsIgnorer) Ignore(path string, directory bool) (IgnoreStatus, bool) {
	_ = "STUB: not implemented"
	// Watch for and ignore any VCS directories.
	return *new(IgnoreStatus), false
}

// Dispatch all other requests to the underlying ignorer.

// IgnoreVCS wraps an ignorer, modifying it to ignore VCS directories.
func IgnoreVCS(ignorer Ignorer) Ignorer { _ = "STUB: not implemented"; return *new(Ignorer) }
