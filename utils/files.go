package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"gowork/log"
)

type myfile struct{}

// File utilities.
var File = myfile{}

// GetFileSize get the length in bytes of file of the specified path.
func (*myfile) GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		log.Error(err)

		return -1
	}

	return fi.Size()
}

// IsExist determines whether the file spcified by the given path is exists.
func (*myfile) IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func (*myfile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}

// IsImg determines whether the specified extension is a image.
func (*myfile) IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// IsDir determines whether the specified path is a directory.
func (*myfile) IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if nil != err {
		log.Warningf("Determines whether [%s] is a directory failed: [%v]", path, err)

		return false
	}

	return fio.IsDir()
}

// CopyFile copies the source file to the dest file.
func (*myfile) CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		if sourceinfo, e := os.Stat(source); nil != e {
			err = os.Chmod(dest, sourceinfo.Mode())

			return
		}
	}

	return nil
}

// CopyDir copies the source directory to the dest directory.
func (*myfile) CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}

	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = File.CopyDir(srcFilePath, destFilePath)
			if err != nil {
				log.Error(err)
			}
		} else {
			err = File.CopyFile(srcFilePath, destFilePath)
			if err != nil {
				log.Error(err)
			}
		}
	}

	return nil
}
