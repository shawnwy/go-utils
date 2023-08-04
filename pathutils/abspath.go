package pathutils

import (
	"os"
	"path/filepath"
	"strings"
)

func AbsPath(inPath string) string {
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = UserHomeDir() + inPath[5:]
	}

	inPath = os.ExpandEnv(inPath)

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}
	return ""
}
