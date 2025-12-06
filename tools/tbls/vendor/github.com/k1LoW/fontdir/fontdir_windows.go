package fontdir

import (
	"os"
	"path/filepath"
)

func getFromUserDirs() []string {
	return []string{}
}

func getFromSystemDirs() []string {
	return []string{
		filepath.Join(os.Getenv("windir"), "Fonts"),
		filepath.Join(os.Getenv("localappdata"), "Microsoft", "Windows", "Fonts"),
	}
}
