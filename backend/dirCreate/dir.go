package dirCreate

import "os"

func EnsureDir(dirName string) error {

	mode := int(0777)
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
