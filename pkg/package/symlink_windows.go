package package

import "os"

func makeSymlink(oldname, newname string) error {
	f, err := os.OpenFile(newname, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.WriteString(oldname)
	return err
}
