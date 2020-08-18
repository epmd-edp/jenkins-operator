package finalizer

import "os"

func ContainsString11(slice []string, s string) {
	FileExists("test")
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
