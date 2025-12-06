package fontdir

import "os"

func Get() []string {
	dirs := getFromUserDirs()
	dirs = append(dirs, getFromSystemDirs()...)
	return dirs
}

func homeDir() string {
	home, _ := os.UserHomeDir()
	return home
}
