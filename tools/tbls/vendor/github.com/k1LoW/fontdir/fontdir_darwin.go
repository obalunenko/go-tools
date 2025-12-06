package fontdir

import (
	"path/filepath"
)

func getFromUserDirs() []string {
	home := homeDir()
	return []string{
		filepath.Join(home, "Library", "Fonts"),
	}
}

func getFromSystemDirs() []string {
	return []string{
		"/Library/Fonts/",
		"/System/Library/Fonts/",
	}
}
