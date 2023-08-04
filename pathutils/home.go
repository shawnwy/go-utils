//go:build !windows

package pathutils

import "os"

func UserHomeDir() string {
	return os.Getenv("HOME")
}
