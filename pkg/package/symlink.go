// +build !windows

package pkg

import "os"

func makeSymlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}
