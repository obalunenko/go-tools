package fontdir

import (
	"os"
	"path/filepath"
	"strings"
)

func getFromUserDirs() []string {
	home := homeDir()
	if dataPath := os.Getenv("XDG_DATA_HOME"); dataPath != "" {
		return []string{filepath.Join(home, ".fonts"), filepath.Join(strings.Replace(dataPath, "~", home, -1), "fonts")}
	}
	return []string{filepath.Join(home, ".fonts"), filepath.Join(home, ".local", "share", "fonts")}
}

func getFromSystemDirs() (dirs []string) {
	home := homeDir()
	if dataPaths := os.Getenv("XDG_DATA_DIRS"); dataPaths != "" {
		for _, dataPath := range filepath.SplitList(dataPaths) {
			dirs = append(dirs, filepath.Join(strings.Replace(dataPath, "~", home, -1), "fonts"))
		}
		return dirs
	}
	return []string{"/usr/local/share/fonts/", "/usr/share/fonts/"}
}
