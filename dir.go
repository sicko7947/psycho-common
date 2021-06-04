package sickocommon

import (
	"os"
	"runtime"
)

func PathCheckAndCreate(s string, perm os.FileMode) error {
	ext, err := PathExists(s)
	if !ext {
		return os.MkdirAll(s, perm)
	}
	return err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAppdataPath() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME") + "\\AppData"
}
