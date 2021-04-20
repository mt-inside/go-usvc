package usvc

import (
	"os/user"
	"path/filepath"
)

// HomePath returns the absolute path to a given name in the current user's home directory
func HomePath(path string) string {
	user, _ := user.Current()
	home := user.HomeDir
	return filepath.Join(home, path)
}
