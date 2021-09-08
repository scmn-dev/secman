package cluster

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func StubBackupCluster() func() {
	orig := BackupClusterFile
	BackupClusterFile = func(_ string) error {
		return nil
	}

	return func() {
		BackupClusterFile = orig
	}
}

func StubWriteCluster(wc io.Writer, wh io.Writer) func() {
	orig := WriteClusterFile
	WriteClusterFile = func(fn string, data []byte) error {
		switch filepath.Base(fn) {
		case "config.yml":
			_, err := wc.Write(data)
			return err
		case "hosts.yml":
			_, err := wh.Write(data)
			return err
		default:
			return fmt.Errorf("write to unstubbed file: %q", fn)
		}
	}

	return func() {
		WriteClusterFile = orig
	}
}

func stubCluster(main, hosts string) func() {
	orig := ReadClusterFile
	ReadClusterFile = func(fn string) ([]byte, error) {
		switch filepath.Base(fn) {
		case "config.yml":
			if main == "" {
				return []byte(nil), os.ErrNotExist
			} else {
				return []byte(main), nil
			}
		case "hosts.yml":
			if hosts == "" {
				return []byte(nil), os.ErrNotExist
			} else {
				return []byte(hosts), nil
			}
		default:
			return []byte(nil), fmt.Errorf("read from unstubbed file: %q", fn)
		}

	}

	return func() {
		ReadClusterFile = orig
	}
}
