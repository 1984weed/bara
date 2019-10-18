package utils

import (
	"os"
)

type FileUtil struct {
	path string
}

func NewFileUtils(path string) *FileUtil {
	return &FileUtil{
		path: path,
	}
}

func (f *FileUtil) Create() error {
	var _, err = os.Stat(f.path)

	if os.IsNotExist(err) {
		var file, err = os.Create(f.path)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}

func (f *FileUtil) Write(content string) error {
	var file, err = os.OpenFile(f.path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (f *FileUtil) WriteCreateFile(content string) error {
	err := f.Create()
	if err != nil {
		return err
	}
	err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}
