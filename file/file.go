package file

import (
	"io"
	"os"
)

func Copy(src string, dest string) (int64, error) {

	sourceFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer newFile.Close()

	n, err := io.Copy(newFile, sourceFile)
	if err != nil {
		return n, err
	}
	return n, nil
}
